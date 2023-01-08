package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Handlers(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all message created by bot itself
	if m.Author.ID == s.State.SessionID {
		return
	}

	// Allow bot to read message
	s.Identify.Intents |= discordgo.IntentMessageContent

	fmt.Println(m.Content)
}
