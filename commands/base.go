package commands

import (
	"fmt"
	
	"github.com/bwmarrin/discordgo"
)

var (
	Commands        = make([]*discordgo.ApplicationCommand, 0)
	CommandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
)

func addCommand(command *discordgo.ApplicationCommand, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	_, exist := CommandHandlers[command.Name]
	if exist {
		panic(fmt.Sprintf("[%s] ← このコマンド名が重複しています！", command.Name))
	}
	// コマンド部分のNameをそのままmapのKeyとして設定しておく
	CommandHandlers[command.Name] = fn
	Commands = append(Commands, command)
}