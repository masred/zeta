package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/masred/zeta/bot"
	"github.com/masred/zeta/config"
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

	discord.AddHandler(bot.MessageCreate)
	discord.AddHandler(bot.MessageFromSlashCommand)

	if err = discord.Open(); err != nil {
		log.Fatalln("Error opening Discord session: ", err)
	}
	defer discord.Close()

	log.Println("Bot ID: ", discord.State.User.ID)
	log.Println("Number of Server: ", len(discord.State.Guilds))
	log.Printf("Logged in as: %v#%v", discord.State.User.Username, discord.State.User.Discriminator)

	appID := &discord.State.User.ID
	commandPrefix := viper.GetString("app.command")
	command := discordgo.ApplicationCommand{
		Name:        commandPrefix,
		Description: "hi👋, aku Zeta😁",
		Options: []*discordgo.ApplicationCommandOption{{
			Name:        "claim-role",
			Description: "Choose role and add emoji",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{{
				Name:        "role",
				Description: "Choose role",
				Type:        discordgo.ApplicationCommandOptionRole,
				Required:    true,
			}, {
				Name:        "emoji",
				Description: "Press \"Win + .\" to add emoji",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			}},
		}},
	}

	registeredCommand, err := discord.ApplicationCommandCreate(*appID, "", &command)
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
