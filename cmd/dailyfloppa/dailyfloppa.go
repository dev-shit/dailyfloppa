package main

import (
    "os"
    "errors"

	"github.com/dev-shit/dailyfloppa/internals/bot"
	"github.com/dev-shit/dailyfloppa/internals/logger"
)

func main() {
    logger.Log("Bot started")
    s := bot.Init()
    s.Open()

    channelId, ok := os.LookupEnv("CHANNEL_ID")
    if !ok {
        panic(errors.New("No channel ID provided"))
    }

    _, err := s.ChannelMessageSend(channelId, "bruh moment")
    if err != nil {
        logger.Log(err.Error())
    }

    s.Close()
}
