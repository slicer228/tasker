package clock

// custom type for abstraction
type Time float64

// interface for working with timestamps
type TimeManager interface {
	GetCurrentTime() Time
}

// interface for comfortable formatting time in specified timezone
type TimeFormatter interface {
	GetFormattedTime(Time) string
	GetFormattedTimeDelta(Time, Time) string
}
