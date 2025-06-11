package app

type AbstractApp interface {
	MustRun()
	Stop()
}
