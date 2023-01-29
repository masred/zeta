package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/masred/zeta/cmd"
	"github.com/masred/zeta/handler"
	"github.com/masred/zeta/pkg/config"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitDefaultConfig(); err != nil {
		log.Fatalln("Error initiating config: ", err.Error())
	}

	discord, err := discordgo.New("Bot " + viper.GetString("app.token"))
	if err != nil {
		log.Fatalln("Error creating Discord session: ", err)
	}

	discord.AddHandler(handler.SetRoleByReactMessageHandler)

	if err = discord.Open(); err != nil {
		log.Fatalln("Error opening Discord session: ", err)
	}
	defer discord.Close()

	log.Printf("Logged in as: %v#%v", discord.State.User.Username, discord.State.User.Discriminator)
	log.Printf("Used in %d servers", len(discord.State.Guilds))

	appID := &discord.State.User.ID
	registeredCommand, err := discord.ApplicationCommandCreate(*appID, "", &cmd.SetRoleCommand)
	if err != nil {
		log.Fatalln("Error creating application command: ", err.Error())
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Bot is running. Press CTRL-C to exit.")
	<-stop
	fmt.Println("")

	if err = discord.ApplicationCommandDelete(*appID, "", registeredCommand.ID); err != nil {
		log.Fatalln("Error deleting application command: ", err.Error())
	}
	log.Println("Successfully deleted application command: ", registeredCommand.Name)
}
