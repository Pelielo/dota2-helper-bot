package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pelielo/dota2-helper-bot/src/helper/actions"
)

func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate, commands []string) {
	switch true {

	// Show documentation on all commands
	case commands[0] == "help":
		s.ChannelMessageSend(m.ChannelID, actions.ShowCommands())

	// Tosses a coin
	case commands[0] == "toss":
		s.ChannelMessageSend(m.ChannelID, actions.CoinToss())

	// Rolls a number. Defaults to 0-100
	case commands[0] == "roll":
		s.ChannelMessageSend(m.ChannelID, actions.RollNumber())

	// Show documentation on lobby command
	case commands[0] == "lobby" && len(commands) == 1:
		s.ChannelMessageSend(m.ChannelID, "Usage: `-lobby player1 player2 player3 ...`")
	// Show documentation on lobby-roles command
	case commands[0] == "lobby-roles" && len(commands) == 1:
		s.ChannelMessageSend(m.ChannelID, "Usage: `-lobby-roles player1 player2 player3 ...`")
	// Error when number of players is not 10
	case (commands[0] == "lobby" || commands[0] == "lobby-roles") && len(commands[1:]) < 10:
		s.ChannelMessageSend(m.ChannelID, "WE NEED MOAR PLAYERS!!! <:unamused_peli:731992316364980286>")
	// Error when number of players is not 10
	case (commands[0] == "lobby" || commands[0] == "lobby-roles") && len(commands[1:]) > 12:
		s.ChannelMessageSend(m.ChannelID, "WE NEED LESS PLAYERS!!! <:sad_carol:731977792471957554>")
	// Randomizes a lobby of 10, 11 or 12 people
	case commands[0] == "lobby" && len(commands[1:]) >= 10 && len(commands[1:]) <= 12:
		s.ChannelMessageSend(m.ChannelID, actions.BuildLobby(commands[1:], false))
	// Randomizes a lobby with roles of 10, 11 or 12 people
	case commands[0] == "lobby-roles" && len(commands[1:]) >= 10 && len(commands[1:]) <= 12:
		s.ChannelMessageSend(m.ChannelID, actions.BuildLobby(commands[1:], true))

}
