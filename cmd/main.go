package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alitdarmaputra/nadeshiko-bot/pkg/handlers"
	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating discord session:", err)
		return
	}

	dg.AddHandler(handlers.Handlers)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening discord session:", err)
		return
	}

	// Wait until CTRL-C or other term signal is received
	fmt.Println("Nadeshiko BOT is running...")
	fmt.Println("Press CTRL-c to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close discord sessions
	dg.Close()
}
