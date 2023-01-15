package love

import (
	"fmt"
	"strings"

	love "github.com/alitdarmaputra/nadeshiko-bot/internal/love/services"
	"github.com/bwmarrin/discordgo"
)

type LoveHandler struct {
}

func NewLoveHandler() *LoveHandler {
	return &LoveHandler{}
}

func (l *LoveHandler) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all message created by bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Allow bot to read message
	s.Identify.Intents |= discordgo.IntentMessageContent

	if strings.HasPrefix(m.Content, "!lovecalc") {
		// Split argument
		args := strings.Split(m.Content, " ")

		if len(args) == 3 {
			var loveService = love.NewLoveService()

			content := loveService.LoveCalc(args[1], args[2])
			_, err := s.ChannelMessageSend(m.ChannelID, content)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		_, err := s.ChannelMessageSend(
			m.ChannelID,
			"Not enough argument. Please provide **two name**",
		)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
