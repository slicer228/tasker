package task

import (
	"context"
	"log/slog"
	"sync"
	"tasker/internal/service/lib/clock"
)

const (
	Ready     = "created"
	Running   = "running"
	Completed = "completed"
)

type Result struct {
	Value any
	err   error
}

type ToCall struct {
	call   func(ctx *context.Context, args ...any) (any, error)
	args   []any
	cancel context.CancelFunc
}

type TaskInfo struct {
	CreatedAt   string
	StartedAt   string
	CompletedAt string
	TimeSpent   string
	Status      string
	Result      *Result
}

type Task struct {
	Runnable
	Statusable
	toCall *ToCall
	result chan *Result
	status string
	log    *slog.Logger
	clock  *clock.Clock
	mu     sync.Mutex
	Timestamps
}

type Timestamps struct {
	createdAt   clock.Time
	startedAt   clock.Time
	completedAt clock.Time
}

func (t *Task) GetInfo() *TaskInfo {
	t.mu.Lock()
	defer t.mu.Unlock()

	info := &TaskInfo{}
	select {
	case res := <-t.result:
		info.Result = res
	default:

	}
	info.CreatedAt = t.clock.GetFormattedTime(t.createdAt)
	info.StartedAt = t.clock.GetFormattedTime(t.startedAt)

	if t.status == Completed {
		info.TimeSpent = t.clock.GetFormattedTimeDelta(t.createdAt, t.completedAt)
		info.CompletedAt = t.clock.GetFormattedTime(t.completedAt)
	} else {
		info.TimeSpent = t.clock.GetFormattedTimeDelta(t.createdAt, t.clock.GetCurrentTime())
	}

	info.Status = t.status

	return info
}

func (t *Task) Run() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.status != Ready {
		return
	}

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		t.toCall.cancel = cancel
		t.startedAt = t.clock.GetCurrentTime()
		res, err := t.toCall.call(&ctx, t.toCall.args...)
		t.completedAt = t.clock.GetCurrentTime()
		t.result <- &Result{res, err}
		t.status = Completed
	}()

	t.status = Running
	t.log.Info("Task started")
}

func (t *Task) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.status != Running {
		return
	}
	t.toCall.cancel()
	t.status = Ready
	t.log.Info("Task stopped")
}

func NewTask(log *slog.Logger, c *clock.Clock, toCall func(ctx *context.Context, args ...any) (any, error), args ...any) *Task {
	t := &Task{
		Timestamps: Timestamps{
			createdAt: c.GetCurrentTime(),
		},
		toCall: &ToCall{
			call: toCall,
			args: args,
		},
		result: make(chan *Result),
		log:    log,
		clock:  c,
		mu:     sync.Mutex{},
		status: Ready,
	}

	log.Info("Task created")

	return t
}
