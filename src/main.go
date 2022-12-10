package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
)

var conf tomlConfig

func init() {
	if _, err := toml.DecodeFile("config/config.toml", &conf); err != nil {
		log.Fatalf("error: could not parse configuration :%v\n", err)
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

	switch {

	case mSplit[0] == "ping":
		ping(s, m)

	}
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Collect current timestamp for comparision
	now := time.Now()

	// Send the message to Discord
	msg, _ := s.ChannelMessageSend(m.ChannelID, "Pong!")

	// Calculate the timestamp from the snowflake
	timestamp, _ := discordgo.SnowflakeTimestamp(msg.ID)

	// Find the difference in the API timestamp and local timestamp
	diff := int32(timestamp.Sub(now).Seconds() * 1000)

	// Send the ping message into the respective channel
	msg, _ = s.ChannelMessageEdit(msg.ChannelID, msg.ID, fmt.Sprintf(":ping_pong: WS Roundtrip: %dms!", diff))
}
