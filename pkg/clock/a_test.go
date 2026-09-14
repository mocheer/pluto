package clock_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/clock"
	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	c := clock.MustParse("2021-06-16 10:09:01", clock.FmtFullDate)
	assert.Equal(t, c.Fmt(clock.FmtDate), "2021-06-16")
	assert.Equal(t, c.Fmt(clock.FmtFullDate), "2021-06-16 10:09:01")
	assert.Equal(t, c.Fmt(clock.FmtCompactDate), "20210616")
	assert.Equal(t, c.Fmt(clock.FmtCompactFullDate), "20210616100901")
}

func TestClock2(t *testing.T) {
	now, _ := json.Marshal(time.Now())
	t.Log(string(now))
	t.Log(clock.New(time.Now()).Fmt(clock.FmtFullDate))

}

func TestClock3(t *testing.T) {
	now, _ := json.Marshal(time.Now().UTC())
	t.Log(string(now))
	t.Log(clock.New(time.Now().UTC()).Fmt(clock.FmtFullDate))

}
