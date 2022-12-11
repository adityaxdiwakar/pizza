package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
)

var conf tomlConfig
var db database

func init() {
	if _, err := toml.DecodeFile("config/config.toml", &conf); err != nil {
		log.Fatalf("error: could not parse configuration :%v\n", err)
	}

	if conf.DiscordConfig.Env == "dev" {
		db = testConnection{}
	} else if conf.DiscordConfig.Env == "prod" {
		// TODO: replace with proper dependency injection
		db = testConnection{}
	}
}

func main() {
	dg, err := discordgo.New(fmt.Sprintf("Bot %s", conf.DiscordConfig.Token))
	if err != nil {
		log.Fatalf("Error creating Discord session due to %v\n", err)
	}

	dg.AddHandler(messageCreate)

	if err = dg.Open(); err != nil {
		log.Fatalf("Error opening connection to Discord: %v\n", err)
	}

	fmt.Println("Pizza is now loaded.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc // block until signal is received

	fmt.Println("Interrupt received, terminating Pizza.")
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, conf.DiscordConfig.Prefix) {
		return
	}

	m.Content = m.Content[len(conf.DiscordConfig.Prefix):]
	mSplit := strings.Split(m.Content, " ")

	// command selector
	switch mSplit[0] {

	case "ping":
		ping(s, m)

	case "userinfo":
		getUserInfo(s, m)
	}
}
