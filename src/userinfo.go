package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
)

func checkValidId(id string) bool {
	for _, c := range id {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func boolToEmoji(b bool) string {
	if b {
		return ":green_circle:"
	}
	return ":red_circle:"
}

func getUserInfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	mArgs := strings.Fields(m.Content)
	if len(mArgs) < 2 {
		_, _ = s.ChannelMessageSend(
			m.ChannelID,
			":warning: No User ID argument passed in.",
		)
		return
	}

	userId := mArgs[1]
	// check the userId passed in is valid
	if len(userId) < 16 && !checkValidId(userId) {
		_, _ = s.ChannelMessageSend(
			m.ChannelID,
			":warning: User ID is not valid.",
		)
		return
	}

	msg, _ := s.ChannelMessageSend(
		m.ChannelID,
		":hourglass: Please wait for querying to complete.",
	)

	content := strings.Builder{}
	username := ""
	for k, v := range db.getRegisteredGuilds() {
		guildResident, err := s.GuildMember(k, userId)

		if guildResident != nil && err == nil {
			username = fmt.Sprintf(
				"%s#%s (<@%s>)",
				guildResident.User.Username,
				guildResident.User.Discriminator,
				guildResident.User.ID,
			)
		}

		content.WriteString(
			fmt.Sprintf("%s: %s\n", boolToEmoji(err == nil), v),
		)
	}

	unamedContent := ""
	if username != "" {
		unamedContent = fmt.Sprintf(" for %s", username)
	}

	_, _ = s.ChannelMessageEdit(
		msg.ChannelID,
		msg.ID,
		fmt.Sprintf(
			"Cross Server Information%s:\n\n%s",
			unamedContent,
			content.String(),
		),
	)
}
