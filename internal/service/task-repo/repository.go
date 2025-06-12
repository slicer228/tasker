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
	maxTasks uint64                //max tasks stored in time
	tasks    map[uint64]*task.Task //map of stored tasks
	mu       sync.RWMutex          //mutex for sync work with map
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

func (t *TaskFarm) StopTask(taskId uint64) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	found := t.tasks[taskId]
	if found == nil {
		return fmt.Errorf("task not found")
	}
	t.tasks[taskId].Stop()

	return nil
}

func (t *TaskFarm) RemoveTask(taskId uint64) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	found := t.tasks[taskId]
	if found == nil {
		return fmt.Errorf("task not found")
	}
	t.tasks[taskId].Stop()
	delete(t.tasks, taskId)

	t.log.Info("Task removed", "task_id", taskId)

	return nil
}

func (t *TaskFarm) RunTask(taskId uint64) error {
	t.mu.RLock()
	defer t.mu.RUnlock()

	found := t.tasks[taskId]
	if found == nil {
		return fmt.Errorf("task not found")
	}
	t.tasks[taskId].Run()

	return nil
}

func (t *TaskFarm) GetTaskInfo(taskId uint64) (*task.TaskFormatted, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	found := t.tasks[taskId]
	if found == nil {
		return nil, fmt.Errorf("task not found")
	}

	return t.tasks[taskId].GetInfoFormatted(), nil
}

func NewTaskFarm(log *slog.Logger, maxTasks uint64) *TaskFarm {
	return &TaskFarm{
		log:      log,
		maxTasks: maxTasks,
		tasks:    make(map[uint64]*task.Task),
		mu:       sync.RWMutex{},
	}
}
