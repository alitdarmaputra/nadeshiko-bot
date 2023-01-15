package help

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type HelpService struct {
}

func NewHelpService() *HelpService {
	return &HelpService{}
}

func (h *HelpService) HelpService(m *discordgo.MessageCreate) string {
	var content = fmt.Sprintf("Hello %s, nice to meet you ^_^\n**Keywords: h, stalk, lovecalc, tod**\n\nFor help type:\n!usage <available keyword>", m.Author.Username)
	return content
}

func (h *HelpService) GetHelp(command string) string {
	if command == "lovecalc" {
		return "Type:\n!lovecalc <name1> <name2>\n**Ex: !lovecalc Nadeshiko Kagamihara**"
	} else if command == "stalk" {
		return "Type:\n!stalk <instagram_username>\n**to show 3 recent photo from instagram**"
	} else if command == "h" {
		return "Type:\n!h\n**to show help**"
	} else if command == "usage" {
		return "Type:\n!usage <command_name>\n**to show how to use a command**"
	} else if command == "tod" {
		return "Type:\n!tod <truth | dare>\n**To get a truth or dare question**"
	} else {
		return "Sorry, keyword not found"
	}
}
