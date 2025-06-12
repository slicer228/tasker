package task_repo

import (
	"context"
	"tasker/internal/service/lib/clock"
	"tasker/internal/service/task-repo/task"
)

// interface to work with repo
type TaskRepo interface {
	CreateTask(c *clock.Clock, toCall func(ctx *context.Context, args ...any) (any, error), args ...any) (uint64, error)
	StopTask(taskId uint64)
	RemoveTask(taskId uint64)
	RunTask(taskId uint64)
	GetTaskInfo(taskId uint64) *task.TaskFormatted
}
