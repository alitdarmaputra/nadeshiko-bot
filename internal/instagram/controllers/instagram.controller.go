package instagram

import (
	"fmt"
	"strings"

	instagram "github.com/alitdarmaputra/nadeshiko-bot/internal/instagram/services"
	"github.com/bwmarrin/discordgo"
)

type InstagramHandler struct {
}

func NewInstagramHandler() *InstagramHandler {
	return &InstagramHandler{}
}

func (i *InstagramHandler) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all message created by bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Allow bot to read message
	s.Identify.Intents |= discordgo.IntentMessageContent

	if strings.HasPrefix(m.Content, "!stalk") {
		var instagramService = instagram.NewInstagramService()

		// Split argument
		args := strings.Split(m.Content, " ")

		// Check if arg is provided
		if len(args) == 2 {
			instagram, err := instagramService.GetUserId(args[1])

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			userFeeds, err := instagramService.GetUserFeeds(instagram.UserID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			instagram.UserFeeds = userFeeds

			for i := 0; i < len(instagram.UserFeeds) && i < 3; i++ {
				if i == 0 {
					_, err = s.ChannelMessageSend(
						m.ChannelID,
						fmt.Sprintf("Here what I get from %s", instagram.Username),
					)
					if err != nil {
						fmt.Println("Error:", err)
					}
				}
				_, err = s.ChannelMessageSend(m.ChannelID, instagram.UserFeeds[i])
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
			return
		}

		_, err := s.ChannelMessageSend(
			m.ChannelID,
			"Not enough argument. Please provide **username**",
		)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
