package handlers

import (
	"fmt"
	"strings"

	"github.com/alitdarmaputra/nadeshiko-bot/pkg/services"
	"github.com/bwmarrin/discordgo"
)

func Handlers(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all message created by bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Allow bot to read message
	s.Identify.Intents |= discordgo.IntentMessageContent

	if strings.HasPrefix(m.Content, "!") {
		if strings.HasPrefix(m.Content, "!usage") {
			// Split argument
			args := strings.Split(m.Content, " ")

			if len(args) == 2 {
				content := services.GetHelp(args[1])
				_, err := s.ChannelMessageSend(m.ChannelID, content)
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			_, err := s.ChannelMessageSend(
				m.ChannelID,
				"Not enough argument. Please provide command name",
			)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if m.Content == "!h" {
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
		} else if strings.HasPrefix(m.Content, "!lovecalc") {
			// Split argument
			args := strings.Split(m.Content, " ")

			if len(args) == 3 {
				content := services.LoveCalc(args[1], args[2])
				_, err := s.ChannelMessageSend(m.ChannelID, content)
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			_, err := s.ChannelMessageSend(m.ChannelID, "Not enough argument. Please provide **two name**")
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if strings.HasPrefix(m.Content, "!tod") {
			// Split argument
			args := strings.Split(m.Content, " ")

			if len(args) == 2 {
				content, err := services.GetTOD(args[1], m.Author.ID)
				if err != nil {
					fmt.Println(err)
					return
				}

				_, err = s.ChannelMessageSend(m.ChannelID, content)
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			_, err := s.ChannelMessageSend(m.ChannelID, "Not enough argument. Please provide **truth or dare word**")
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
