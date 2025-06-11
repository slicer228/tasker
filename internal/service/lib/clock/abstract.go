package clock

type Time float64

type TimeManager interface {
	GetCurrentTime() Time
}

type TimeFormatter interface {
	GetFormattedTime(Time) string
	GetFormattedTimeDelta(Time, Time) string
}
