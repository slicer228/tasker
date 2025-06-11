package clock

import (
	"testing"
	"time"
)

func TestFormatting(t *testing.T) {
	timer := NewTimer(time.UTC)
	var res string

	res = timer.GetFormattedTime(0)
	if res != "1970-01-01 00:00:00.00" {
		t.Errorf("clock.GetFormattedTime(0) returned %s\nExpected: 1970-01-01 00:00:00.00", res)
	}
	res = timer.GetFormattedTimeDelta(0, 0)
	if res != "00:00:00.00" {
		t.Errorf("clock.GetFormattedTimeDelta(0, 0) returned %s\nExpected: 00:00:00.00", res)
	}
	res = timer.GetFormattedTimeDelta(0, 5)
	if res != "00:00:05.00" {
		t.Errorf("clock.GetFormattedTimeDelta(0, 5) returned %s\nExpected: 00:00:05.00", res)
	}
	res = timer.GetFormattedTimeDelta(0, 5.1324)
	if res != "00:00:05.13" {
		t.Errorf("clock.GetFormattedTimeDelta(0, 5.1324) returned %s\nExpected: 00:00:05.13", res)
	}
	res = timer.GetFormattedTimeDelta(0, 5.1399)
	if res != "00:00:05.13" {
		t.Errorf("clock.GetFormattedTimeDelta(0, 5.1399) returned %s\nExpected: 1970-01-01 00:00:05.13", res)
	}
	res = timer.GetFormattedTimeDelta(timer.GetCurrentTime(), timer.GetCurrentTime()+5)
	if res != "00:00:05.00" {
		t.Errorf("clock.GetFormattedTimeDelta(clock.GetCurrentTime(), clock.GetCurrentTime() + 5) returned %s\nExpected: 00:00:05.00", res)
	}

}
