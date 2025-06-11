package task

type Runnable interface {
	Run() error
	Stop() error
}

type Statusable interface {
	GetInfo() *TaskInfo
}
