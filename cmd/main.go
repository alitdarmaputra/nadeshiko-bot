package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	help "github.com/alitdarmaputra/nadeshiko-bot/internal/help/controllers"
	instagram "github.com/alitdarmaputra/nadeshiko-bot/internal/instagram/controllers"
	love "github.com/alitdarmaputra/nadeshiko-bot/internal/love/controllers"
	tod "github.com/alitdarmaputra/nadeshiko-bot/internal/tod/controllers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error reading .env file")
	}

	Token = os.Getenv("BOT_TOKEN")
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating discord session:", err)
		return
	}

	dg.AddHandler(instagram.NewInstagramHandler().Handler)
	dg.AddHandler(love.NewLoveHandler().Handler)
	dg.AddHandler(tod.NewTodhandler().Handler)
	dg.AddHandler(help.NewHelpHandler().Handler)

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
