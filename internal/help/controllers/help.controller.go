package help

import (
	"fmt"
	"strings"

	help "github.com/alitdarmaputra/nadeshiko-bot/internal/help/services"
	"github.com/bwmarrin/discordgo"
)

type HelpHandler struct {
}

func NewHelpHandler() *HelpHandler {
	return &HelpHandler{}
}

func (h *HelpHandler) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var helpService = help.NewHelpService()

	if m.Content == "!h" {
		var content string = helpService.HelpService(m)
		_, err := s.ChannelMessageSend(m.ChannelID, content)
		if err != nil {
			fmt.Println("Error:", err)
		}
		return
	}

	if strings.HasPrefix(m.Content, "!usage") {
		// Split argument
		args := strings.Split(m.Content, " ")

		if len(args) == 2 {
			content := helpService.GetHelp(args[1])
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
	}
}
