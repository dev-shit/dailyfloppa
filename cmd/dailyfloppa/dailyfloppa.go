package main

import (
	"github.com/dev-shit/dailyfloppa/internals/bot"
)

func main() {
    s := bot.Init()

    bot.SendFloppa(s)

    s.Close()
}
