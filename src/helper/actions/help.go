package actions

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ShowCommands() *discordgo.MessageEmbed {
	commands := []string{
		"`-toss`: Tosses a coin and outputs heads or tails.",
		"`-roll`: Randomly chooses a value between two numbers. Defaults to 0-100.",
		"`-lobby`: Creates a lobby with 10-12 randomly chosen players.",
		"`-lobby-roles`: Creates a lobby with 10-12 randomly chosen players and assigns each of them a role.",
	}

	return &discordgo.MessageEmbed{
		Title:       "Commands available",
		Color:       0xbf2121, // red
		Description: strings.Join(commands, "\n"),
	}
}
