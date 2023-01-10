package handlers

import (
	"fmt"
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
		if m.Content == "!help" {
			var content string = services.HelpService(m)
			_, err := s.ChannelMessageSend(m.ChannelID, content)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if strings.HasPrefix(m.Content, "!stalk") {
			// Split argument
			args := strings.Split(m.Content, " ")

			// Check if arg is provided
			if len(args) == 2 {
				userInfo, err := services.GetUserId(args[1])
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				userFeeds, err := services.GetUserFeeds(userInfo.UserID)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				userInfo.UserFeeds = userFeeds
				fmt.Println("List url", userInfo.UserFeeds)
				for i := 0; i < len(userInfo.UserFeeds) && i < 3; i++ {
					if i == 0 {
						_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Here what I get from %s", userInfo.Username))
						if err != nil {
							fmt.Println("Error:", err)
						}
					}
					_, err = s.ChannelMessageSend(m.ChannelID, userInfo.UserFeeds[i])
					if err != nil {
						fmt.Println("Error:", err)
					}
				}
				return
			}

			_, err := s.ChannelMessageSend(m.ChannelID, "Not enough argument. Please provide **username**")
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			var content string = services.NotFoundService()
			_, err := s.ChannelMessageSend(m.ChannelID, content)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}
