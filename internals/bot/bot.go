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

    return s
}

func SendFloppa(s *discordgo.Session) {
    channelId, ok := os.LookupEnv("CHANNEL_ID")
    if !ok {
        panic(errors.New("No channel ID provided"))
    }

    _, err := s.ChannelMessageSendComplex(
        channelId,
        &discordgo.MessageSend{
            Embed: &discordgo.MessageEmbed{
                Title: "Floppa of the day",
                Image: &discordgo.MessageEmbedImage{
                    URL: "https://media.tenor.co/videos/a97a96e86b5805aec031c6625b1e13c0/mp4",
                    Width: 300,
                    Height: 300,
                },
            },
        },
    )

    if err != nil {
        logger.Log(err.Error())
    }

    logger.Log("Floppa sent")
}

