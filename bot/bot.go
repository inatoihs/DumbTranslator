package bot

import (
	"DumbTranslator/Convert"
	"DumbTranslator/config"
	"DumbTranslator/commands"
	"log"

	"github.com/bwmarrin/discordgo"
)

var ID string
var goBot *discordgo.Session
var GuildID string

var (
	commandList     = commands.Commands
	commandHandlers = commands.CommandHandlers
)

func Run() {
	// create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
		return
	}
	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}
	ID = user.ID
	goBot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		return
	}
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commandList))
	for i, v := range commandList {
		cmd, err := goBot.ApplicationCommandCreate(goBot.State.User.ID, GuildID, v)
		if err != nil {
			goBot.Close()
		}
		registeredCommands[i] = cmd
	}
	defer goBot.Close()
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == ID {
		return
	}
	// If the message is "Hi" reply with "Hi Back!!"
	if m.Content != "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, Convert.ConvertKorosuke(m.Content))
	}
}