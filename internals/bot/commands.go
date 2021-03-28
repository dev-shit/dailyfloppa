package bot

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dev-shit/dailyfloppa/internals/logger"
)

type Command struct {
    Name    string
    Handler CommandHandler
}

type CommandHandler func(s *discordgo.Session, m *discordgo.MessageCreate)

var commands = []*Command{
    {
        Name: "floppa",
        Handler: floppaHandler,
    },
    {
        Name: "bruhmoment",
        Handler: bruhmomentHandler,
    },
}

func floppaHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    SendFloppa(s)
}

func bruhmomentHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    s.ChannelMessageSend(m.ChannelID, "bruh moment")
}

func RegisterCommands(s *discordgo.Session) {
    s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
        cmd := strings.Split(m.Content, " ")[0]
        if !strings.HasPrefix(cmd, "f!") {
            return
        }

        for _, c := range commands {
            if cmd == "f!" + c.Name {
                logger.Log(
                    fmt.Sprintf("Command '%s' issued by @%s", cmd, m.Author.Username),
                )
                c.Handler(s, m)
                return
            }
        }

        emoji, _ := os.LookupEnv("NOT_FOUND_EMOJI")
        err := s.MessageReactionAdd(m.ChannelID, m.ID, emoji)
        if err != nil {
            logger.Log(err.Error())
        }
    })
}

