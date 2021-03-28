package logger

import (
	"fmt"
	"time"
)

func Log(msg string) {
    fmt.Printf(
        "\033[37m[%s] [DF] \033[0m%s\n",
        time.Now().Format(time.Stamp),
        msg,
    )
}
