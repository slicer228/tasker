package task

type Runnable interface {
	Run() error
	Stop() error
}

type Statusable interface {
	GetInfoFormatted() *TaskFormatted
}

type Callable interface {
	Call() (any, error)
	Cancel() error
}
