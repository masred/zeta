package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func SetRoleByReactMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
						Title:       "Klaim Role Kamu",
						Description: "React pesan ini untuk mendapatkan role ya~",
						Color:       0xe65c93, // red color
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   role.Name,
								Value:  emoji,
								Inline: true,
							},
						},
					}},
				},
			},
		); err != nil {
			log.Println("Error adding handler application command: ", err.Error())
		}
		message, err := s.InteractionResponse(i.Interaction)
		if err != nil {
			log.Fatalln("Error getting interaction response: ", err.Error())
		}
		s.MessageReactionAdd(message.ChannelID, message.ID, emoji)
		s.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
			if r.Emoji.Name == emoji {
				if r.UserID == s.State.User.ID {
					return
				}
				s.GuildMemberRoleAdd(r.GuildID, r.UserID, role.ID)
			}
		})
		s.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
			if r.Emoji.Name == emoji {
				if r.UserID == s.State.User.ID {
					return
				}
				s.GuildMemberRoleRemove(r.GuildID, r.UserID, role.ID)
			}
		})
	}
}
