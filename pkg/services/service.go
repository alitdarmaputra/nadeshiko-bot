package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HelpService(s *discordgo.Session, m *discordgo.MessageCreate) {
	var content = fmt.Sprintf("Hello %s, nice to meet you ^_^\n**Keywords: help**\n\nFor help type:\n!usage <available keyword>", m.Author.Username)
	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func NotFoundService(s *discordgo.Session, m *discordgo.MessageCreate) {
	var content = fmt.Sprintf(
		"Keyword not found " +
			"\n" +
			"\nTry **!help** to get list keywords",
	)

	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
