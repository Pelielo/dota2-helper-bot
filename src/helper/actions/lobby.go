package actions

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func BuildLobby(players []string, add_roles bool) *discordgo.MessageEmbed {
	randomPlayers := shuffleArray(players)
	roles := []string{
		"Off lane",
		"Hard support",
		"Soft support",
		"Safe lane",
		"Mid lane",
		"Coach",
	}

	var radiantPlayers []string
	var direPlayers []string

	if len(randomPlayers)%2 == 0 {
		radiantPlayers = randomPlayers[:len(players)/2]
		direPlayers = randomPlayers[len(players)/2:]
	} else {
		leftover := rand.Intn(2)
		radiantPlayers = randomPlayers[:len(players)/2+leftover]
		direPlayers = randomPlayers[len(players)/2+leftover:]
	}

	if add_roles {
		for player := 0; player < len(radiantPlayers); player++ {
			radiantPlayers[player] = roles[player] + " - " + radiantPlayers[player]
		}
		for player := 0; player < len(direPlayers); player++ {
			direPlayers[player] = roles[player] + " - " + direPlayers[player]
		}
	}
	return buildLobbyMsg(radiantPlayers, direPlayers)
}

func shuffleArray(a []string) []string {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

func buildLobbyMsg(radiantPlayers []string, direPlayers []string) *discordgo.MessageEmbed {

	// description := "**The Radiant**\n" +
	// 	"```\n" + strings.Join(radiantPlayers, "\n") + "\n```" +
	// 	"\n**The Dire**\n" +
	// 	"```\n" + strings.Join(direPlayers, "\n") + "\n```"

	return &discordgo.MessageEmbed{
		Title: "Lobby results",
		Color: 0xbf2121, // red
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "The Radiant",
				Value:  "```\n" + strings.Join(radiantPlayers, "\n") + "\n```",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "The Dire",
				Value:  "```\n" + strings.Join(direPlayers, "\n") + "\n```",
				Inline: true,
			},
		},
	}
}
