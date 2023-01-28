package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "Jangan dong~😳",
		Description: "This is the description for my embed message.",
		Color:       0xe65c93, // red color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Field 1",
				Value:  "Value 1",
				Inline: true,
			},
			{
				Name:   "Field 2",
				Value:  "Value 2",
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "My Embed Footer",
		},
	}
	prefix := viper.GetString("app.prefix")
	switch m.Content {
	case fmt.Sprintf("%srole dong", prefix):
		message, _ := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		s.MessageReactionAdd(m.ChannelID, message.ID, "🔥")
		s.AddHandler(MessageReactionAdd)
		s.AddHandler(MessageReactionRemove)
	}
}

func MessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.Emoji.Name == "🔥" {
		if r.UserID == s.State.User.ID {
			return
		}
		s.GuildMemberRoleAdd(r.GuildID, r.UserID, "732086420356857937")
	}
}

func MessageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	if r.Emoji.Name == "🔥" {
		if r.UserID == s.State.User.ID {
			return
		}
		s.GuildMemberRoleRemove(r.GuildID, r.UserID, "732086420356857937")
	}
}

func MessageFromSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commadPrefix := viper.GetString("app.command")
	data := i.ApplicationCommandData()
	switch data.Name {
	case commadPrefix:
		role := data.Options[0].Options[0].RoleValue(s, i.GuildID)
		emoji := data.Options[0].Options[1].StringValue()
		if err := s.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{{
						Title:       "Hai 👋",
						Description: "React pesan ini untuk mendapatkan role ya 😁",
						Color:       0xe65c93, // red color
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   role.Name,
								Value:  emoji,
								Inline: true,
							},
						},
						Footer: &discordgo.MessageEmbedFooter{
							Text: "My Embed Footer",
						},
					}},
				},
			},
		); err != nil {
			log.Println("Error adding handler application command: ", err.Error())
		}
	}
}
