package handlers

import (
	"strings"

	"github.com/alitdarmaputra/nadeshiko-bot/pkg/services"
	"github.com/bwmarrin/discordgo"
)

func Handlers(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all message created by bot itself
	if m.Author.ID == s.State.SessionID {
		return
	}

	// Allow bot to read message
	s.Identify.Intents |= discordgo.IntentMessageContent

	if isCommand := strings.HasPrefix(m.Content, "!"); isCommand {
		switch m.Content {
		case "!help":
			services.HelpService(s, m)
		default:
			services.NotFoundService(s, m)
		}
	}
}
