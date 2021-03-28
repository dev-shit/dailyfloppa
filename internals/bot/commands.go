package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
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
            if strings.Split(m.Content, " ")[0] == "f!" + c.Name {
                c.Handler(s, m)
            }
        }
    })
}

