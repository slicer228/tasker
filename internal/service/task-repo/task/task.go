package task

import (
	"context"
	"log/slog"
	"sync"
	"tasker/internal/service/lib/clock"
)

const (
	Ready        = "ready"
	Running      = "running"
	JobRunning   = "job_running"
	JobCompleted = "job_completed"
	JobError     = "job_error"
)

type Result struct {
	Value any
	err   error
}

type ToCall struct {
	Callable
	call   func(ctx *context.Context, args ...any) (any, error)
	args   []any
	cancel context.CancelFunc
}

type Task struct {
	Runnable
	Statusable
	createdAt clock.Time
	toCall    *ToCall //function to proceed
	jobs      []*Job  //storage for history of task jobs
	status    string
	log       *slog.Logger
	clock     *clock.Clock
	mu        sync.Mutex //mutex for sync
}

type Job struct {
	Timestamps
	status    string
	JobNumber int     //number in one task
	Result    *Result //contains result of toCall obj
}

type Timestamps struct {
	startedAt   clock.Time
	completedAt clock.Time
	timeSpent   clock.Time
}

type TaskFormatted struct {
	Status    string          `json:"status"`
	CreatedAt string          `json:"createdAt"`
	Jobs      []*JobFormatted `json:"jobs"`
}

type JobFormatted struct {
	TimestampsFormatted
	Status    string  `json:"status"`
	JobNumber int     `json:"jobNumber"`
	Result    *Result `json:"result"`
}

type TimestampsFormatted struct {
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
	TimeSpent   string `json:"timeSpent"`
}

func (tc *ToCall) Call() (any, error) {
	ctx, cancel := context.WithCancel(context.Background())
	tc.cancel = cancel
	return tc.call(&ctx, tc.args...)
}

func (tc *ToCall) Cancel() {
	if tc.cancel != nil {
		tc.cancel()
	}
}

func (t *Task) GetInfoFormatted() *TaskFormatted {
	t.mu.Lock()
	defer t.mu.Unlock()

	fjobs := make([]*JobFormatted, len(t.jobs), len(t.jobs))

	for i, v := range t.jobs {
		jform := &JobFormatted{}

		jform.JobNumber = v.JobNumber
		jform.Status = v.status
		jform.Result = v.Result

		if jform.Status != JobRunning {
			jform.CompletedAt = t.clock.GetFormattedTime(v.completedAt)
			jform.TimeSpent = t.clock.GetFormattedTimeDelta(v.startedAt, v.completedAt)
		} else {
			jform.TimeSpent = t.clock.GetFormattedTimeDelta(v.startedAt, t.clock.GetCurrentTime())
		}

		jform.StartedAt = t.clock.GetFormattedTime(v.startedAt)

		fjobs[i] = jform
	}
	return &TaskFormatted{
		Status:    t.status,
		CreatedAt: t.clock.GetFormattedTime(t.createdAt),
		Jobs:      fjobs,
	}
}

func (t *Task) Run() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.status != Ready {
		return
	}

	job := &Job{}
	t.jobs = append(t.jobs, job)
	job.JobNumber = len(t.jobs)

	go func() {

		job.startedAt = t.clock.GetCurrentTime()
		job.status = JobRunning

		res, err := t.toCall.Call()

		job.completedAt = t.clock.GetCurrentTime()
		job.timeSpent = job.completedAt - job.startedAt
		job.Result = &Result{res, err}

		if err != nil {
			job.status = JobError
		} else {
			job.status = JobCompleted
		}

		t.mu.Lock()
		defer t.mu.Unlock()
		t.status = Ready
	}()

	t.status = Running
	t.log.Info("Task started working")
}

func (t *Task) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.status != Running {
		return
	}
	t.toCall.Cancel()
	t.status = Ready
	t.log.Info("Signal stop sended to job")
}

func NewTask(log *slog.Logger, c *clock.Clock, toCall func(ctx *context.Context, args ...any) (any, error), args ...any) *Task {
	t := &Task{
		createdAt: c.GetCurrentTime(),
		toCall: &ToCall{
			call: toCall,
			args: args,
		},
		jobs:   make([]*Job, 0),
		log:    log,
		clock:  c,
		mu:     sync.Mutex{},
		status: Ready,
	}

	log.Info("Task created")

	return t
}
