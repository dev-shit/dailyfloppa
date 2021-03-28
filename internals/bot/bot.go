package bot

import (
	"errors"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Init() *discordgo.Session {
    token, ok := os.LookupEnv("BOT_TOKEN")
    if !ok {
        panic(errors.New("No token provided"))
    }

    s, err := discordgo.New("Bot " + token)
    if err != nil {
        panic(err)
    }

    return s
}
