package task

// interface to define methods of runnable object
type Runnable interface {
	Run()
	Stop()
}

// interface to objects, which contains statuses
type Statusable interface {
	GetInfoFormatted() *TaskFormatted
}

// interface for callable objects
type Callable interface {
	Call() (any, error)
	Cancel()
}
