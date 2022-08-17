package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	// Nameで定義された文字列がKeyになるので同時に書ける
	addCommand(
		&discordgo.ApplicationCommand{
			Name:        "parrot-return",
			Description: "parrot return",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "hi!",
				},
			})
		},
	)
}