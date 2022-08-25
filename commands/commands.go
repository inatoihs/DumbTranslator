package commands

import (
	"DumbTranslator/Convert"

	"github.com/bwmarrin/discordgo"
)

// func onOjaruCommand(s *discordgo.Session, i *discordgo.InteractionCreate){
	
// }

func init() {
	//Nameで定義された文字列がKeyになるので同時に書ける
	addCommand(
		&discordgo.ApplicationCommand{
			Name:        "korosuke",
			Description: "コロ助の口調に変換します",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var m discordgo.MessageCreate
			if m.Content != "" {
				_, _ = s.ChannelMessageSend(m.ChannelID, Convert.ConvertKorosuke(m.Content))
			}
			if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "hi!",
				},
			}); err != nil {
				return
			}
		},
	)
	addCommand(
		&discordgo.ApplicationCommand{
			Name:        "ojaru",
			Description: "おじゃる丸の口調に変換します",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var m discordgo.MessageCreate
			if m.Content != "" {
				_, _ = s.ChannelMessageSend(m.ChannelID, Convert.ConvertOjaru(m.Content))
			}
			if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "hi!",
				},
			}); err != nil {
				return
			}
		},
	)
}