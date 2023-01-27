package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/masred/zeta/bot"
	"github.com/masred/zeta/config"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitDefaultConfig(); err != nil {
		log.Println(err.Error())
	}

	dg, err := discordgo.New("Bot " + viper.GetString("token"))
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(bot.MessageCreate)

	if err = dg.Open(); err != nil {
		log.Println("Error opening Discord session: ", err)
	}
	defer dg.Close()

	fmt.Println("Bot is running. Press CTRL-C to exit.")
	<-make(chan struct{})
}
