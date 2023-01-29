package cmd

import (
	"github.com/bwmarrin/discordgo"
)

var (
	SetRoleCommand = discordgo.ApplicationCommand{
		Name:        "zeta",
		Description: "hi👋, aku Zeta😁",
		Options: []*discordgo.ApplicationCommandOption{{
			Name:        "set-role-claim",
			Description: "Set role claim by reacting the message",
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
)
