package task_repo

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"tasker/internal/service/lib/clock"
	"tasker/internal/service/task-repo/task"
)

var tasksCount uint64 = 0

type TaskFarm struct {
	TaskRepo
	log      *slog.Logger
	maxTasks uint64
	tasks    map[uint64]*task.Task
	mu       sync.RWMutex
}

func New(log *slog.Logger, maxTasks uint64) *TaskFarm {
	return &TaskFarm{
		log:      log,
		maxTasks: maxTasks,
	}
}

func (t *TaskFarm) CreateTask(c *clock.Clock, toCall func(ctx *context.Context, args ...any) (any, error), args ...any) (uint64, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.maxTasks == tasksCount {
		return 0, fmt.Errorf("max tasks reached")
	}

	tasksCount++
	taskId := tasksCount

	task := task.NewTask(t.log.With("task_id", taskId), c, toCall, args...)

	t.tasks[taskId] = task

	return taskId, nil
}

func (t *TaskFarm) GetTask(taskId uint64) *task.Task {
	return t.tasks[taskId]
}

func (t *TaskFarm) StopTask(taskId uint64) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	t.tasks[taskId].Stop()
}

func (t *TaskFarm) RemoveTask(taskId uint64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.tasks[taskId].Stop()
	delete(t.tasks, taskId)
}

func (t *TaskFarm) RunTask(taskId uint64) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	t.tasks[taskId].Run()
}

func (t *TaskFarm) GetTaskInfo(taskId uint64) *task.TaskInfo {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.tasks[taskId].GetInfo()
}
