package main

import (
	"DumbTranslator/bot"
	"DumbTranslator/config"
	"log"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Run()
	<-make(chan struct{})
	return
}
