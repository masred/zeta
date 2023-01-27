package bot

import (
	"fmt"
	"strings"

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
	if strings.HasPrefix(m.Content, fmt.Sprintf("%srole dong", viper.GetString("prefix"))) {
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
