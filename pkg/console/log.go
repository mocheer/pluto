package console

import (
	"fmt"

	"github.com/mocheer/pluto/pkg/ts/clock"
)

func Log(msg string) {
	fmt.Println(fmt.Sprintf(Green+"[%s] "+Reset+"%s", clock.Now().Fmt(clock.FmtFullDate), msg))
}
