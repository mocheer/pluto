package ts

import (
	"time"

	"github.com/mocheer/pluto/ts/clock"
)

func NewClock(date time.Time) *clock.Clock {
	return clock.New(date)
}
