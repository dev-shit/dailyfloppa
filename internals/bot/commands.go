package bot

import (
	"fmt"
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
}

func floppaHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    SendFloppa(s)
}

func RegisterCommands(s *discordgo.Session) {
    s.AddHandler(func(handlerSession *discordgo.Session, m *discordgo.MessageCreate){
        for _, c := range commands {
            cmd := strings.Split(m.Content, " ")[0]

            if cmd == "f!" + c.Name {
                logger.Log(
                    fmt.Sprintf("Command '%s' issued by @%s", cmd, m.Author.Username),
                )
                c.Handler(s, m)
            }
        }
    })
}

