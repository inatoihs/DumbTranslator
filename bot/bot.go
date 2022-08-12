package bot

import (
	"DumbTranslator/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

var ID string
var goBot *discordgo.Session

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
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		return
	}
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == ID {
		return
	}
	// If the message is "Hi" reply with "Hi Back!!"
	if m.Content != "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, m.Content)
	}
}
