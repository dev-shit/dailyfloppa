package bot

import (
	"errors"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/dev-shit/dailyfloppa/internals/logger"
)

func Init() *discordgo.Session {
    logger.Log("Initiating connection...")

    token, ok := os.LookupEnv("BOT_TOKEN")
    if !ok {
        panic(errors.New("No token provided"))
    }

    s, err := discordgo.New("Bot " + token)
    if err != nil {
        panic(err)
    }

    if err := s.Open(); err != nil {
        panic(err)
    }

    return s
}

func getEmbed() *discordgo.MessageEmbed {
    return &discordgo.MessageEmbed{
        Title: "Floppa of the day",
        Description: "One floppa a day keeps the doctor away",
        URL: "https://github.com/dev-shit/dailyfloppa",
        Image: &discordgo.MessageEmbedImage{
            URL: "https://media1.tenor.com/images/91e86ad9d8c1472d361b554a9e905f08/tenor.gif",
        },
    }
}

func SendFloppa(s *discordgo.Session) {
    channelId, ok := os.LookupEnv("CHANNEL_ID")
    if !ok {
        panic(errors.New("No channel ID provided"))
    }

    _, err := s.ChannelMessageSendComplex(
        channelId,
        &discordgo.MessageSend{Embed: getEmbed()},
    )

    if err != nil {
        logger.Log(err.Error())
        return
    }
}

