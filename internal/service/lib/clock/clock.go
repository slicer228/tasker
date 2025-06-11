package clock

import "time"

var timeFormat string = "2006-01-02 15:04:05.00"
var timeDeltaFormat string = "15:04:05.00"

type Clock struct {
	TimeManager
	TimeFormatter
	timezone *time.Location
}

func (t *Clock) GetCurrentTime() Time {
	return Time(time.Now().UnixNano()) / 1e9
}

func (t *Clock) GetFormattedTime(timestamp Time) string {
	return time.Unix(int64(timestamp), int64((timestamp-Time(int64(timestamp)))*1e9)).In(t.timezone).Format(timeFormat)
}

func (t *Clock) GetFormattedTimeDelta(from Time, to Time) string {
	timestamp := to - from
	return time.Unix(int64(timestamp), int64((timestamp-Time(int64(timestamp)))*1e9)).In(t.timezone).Format(timeDeltaFormat)
}

func NewTimer(timezone *time.Location) *Clock {
	return &Clock{timezone: timezone}
}
