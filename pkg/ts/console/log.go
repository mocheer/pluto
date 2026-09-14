package console

import (
	"fmt"

	"github.com/mocheer/pluto/pkg/clock"
)

func Log(msg string) {
	fmt.Println(fmt.Sprintf(Green+"[%s] "+Reset+"%s", clock.Now().Fmt(clock.FmtFullDate), msg))
}

func Warn(msg string) {
	fmt.Println(fmt.Sprintf(Red+"[%s] "+Reset+"%s", clock.Now().Fmt(clock.FmtFullDate), msg))
}

func Error(msg string) {
	fmt.Println(fmt.Sprintf(Yellow+"[%s] "+Reset+"%s", clock.Now().Fmt(clock.FmtFullDate), msg))
}
