package ts

import (
	"time"

	"github.com/mocheer/pluto/ts/clock"
)

func NewClock(date time.Time) *clock.Clock {
	return clock.New(date)
}

func NewClockFromString(date string, fmt string) *clock.Clock {
	return clock.MustParse(date, fmt)
}
