package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dev-shit/dailyfloppa/internals/bot"
	"github.com/dev-shit/dailyfloppa/internals/logger"
)

func main() {
    s := bot.Init()
    logger.Log("Bot initialized, press CTRL+C to quit")

    // bot.SendFloppa(s)
    bot.RegisterCommands(s)

    sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sigChan
    s.Close()
}
