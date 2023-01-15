package tod

import (
	"fmt"
	"strings"

	tod "github.com/alitdarmaputra/nadeshiko-bot/internal/tod/services"
	"github.com/bwmarrin/discordgo"
)

type TodHandler struct {
}

func NewTodhandler() *TodHandler {
	return &TodHandler{}
}

func (t *TodHandler) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!tod") {
		// Split argument
		args := strings.Split(m.Content, " ")

		if len(args) == 2 {
			var todService = tod.NewTodService()
			content, err := todService.GetTOD(args[1], m.Author.ID)
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

		_, err := s.ChannelMessageSend(
			m.ChannelID,
			"Not enough argument. Please provide **truth or dare word**",
		)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
