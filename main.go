package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var Token string


func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

}

func main() {
	var Token string = "" //discord bot token goes here
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error occurred creating discord session - ", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error occurred opening the connection - ", err)
		return
	}

	fmt.Println("Gary Bot is now live!\nPress CTRL+C to quit the program.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

}
