package clock_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	c := clock.MustParse("2021-06-16 10:09:01", clock.FmtFullDate)
	assert.Equal(t, c.Fmt(clock.FmtDate), "2021-06-16")
	assert.Equal(t, c.Fmt(clock.FmtFullDate), "2021-06-16 10:09:01")
	assert.Equal(t, c.Fmt(clock.FmtCompactDate), "20210616")
	assert.Equal(t, c.Fmt(clock.FmtCompactFullDate), "20210616100901")
}
