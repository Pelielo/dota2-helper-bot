package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token  string
	Prefix string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	flag.StringVar(&Prefix, "p", "-", "Prefix for commands")
	flag.Parse()
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("%v said: %v\n", m.Author.Username, m.Content)

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if msg := m.Content; strings.HasPrefix(msg, Prefix) {
		switch commands := strings.Split(msg[len(Prefix):], " "); {
		// Show documentation on all commands
		case commands[0] == "help":
			s.ChannelMessageSend(m.ChannelID, show_commands())

		// Tosses a coin
		case commands[0] == "toss":
			s.ChannelMessageSend(m.ChannelID, coin_toss())

		// Rolls a number. Defaults to 0-100
		case commands[0] == "roll":
			s.ChannelMessageSend(m.ChannelID, roll_number())

		// Show documentation on lobby command
		case commands[0] == "lobby" && len(commands) == 1:
			s.ChannelMessageSend(m.ChannelID, "Usage: `-lobby player1 player2 player3 ...`")
		// Show documentation on lobby_roles command
		case commands[0] == "lobby_roles" && len(commands) == 1:
			s.ChannelMessageSend(m.ChannelID, "Usage: `-lobby_roles player1 player2 player3 ...`")
		// Error when number of players is not 10
		case (commands[0] == "lobby" || commands[0] == "lobby_roles") && len(commands[1:]) < 10:
			s.ChannelMessageSend(m.ChannelID, "WE NEED MOAR PLAYERS!!! <:unamused_peli:731992316364980286>")
		// Error when number of players is not 10
		case (commands[0] == "lobby" || commands[0] == "lobby_roles") && len(commands[1:]) > 12:
			s.ChannelMessageSend(m.ChannelID, "WE NEED LESS PLAYERS!!!")
		// Randomizes a lobby of 10, 11 or 12 people
		case commands[0] == "lobby" && len(commands[1:]) >= 10 && len(commands[1:]) <= 12:
			s.ChannelMessageSend(m.ChannelID, build_lobby(shuffle_array(commands[1:]), false))
		// Randomizes a lobby with roles of 10, 11 or 12 people
		case commands[0] == "lobby" && len(commands[1:]) >= 10 && len(commands[1:]) <= 12:
			s.ChannelMessageSend(m.ChannelID, build_lobby(shuffle_array(commands[1:]), true))
		}
	}
}

func shuffle_array(a []string) []string {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

func build_lobby(players []string, add_roles bool) string {
	// roles := []string{
	// 	"Offlaner",
	// 	"Hard support",
	// 	"Soft support",
	// 	"Safe lane",
	// 	"Mid lane",
	// }

	var radiant_players []string
	var dire_players []string

	if len(players)%2 == 0 {
		radiant_players = players[:len(players)/2]
		dire_players = players[len(players)/2:]
	} else {
		leftover := rand.Intn(2)
		fmt.Println(leftover)
		radiant_players = players[:len(players)/2+leftover]
		dire_players = players[len(players)/2+leftover:]
	}

	if add_roles {
		return ""
	} else {
		return build_lobby_msg(radiant_players, dire_players)
	}
}

func build_lobby_msg(radiant_players []string, dire_players []string) string {
	// players[:5]
	// players[5:]
	return "**The Radiant**\n" +
		"```\n" + strings.Join(radiant_players, "\n") + "\n```" +
		"\n**The Dire**\n" +
		"```\n" + strings.Join(dire_players, "\n") + "\n```"
}

func coin_toss() string {
	coin := []string{
		"heads",
		"tails",
	}

	return coin[rand.Intn(len(coin))]
}

func show_commands() string {
	commands := []string{
		"`-toss`: Tosses a coin and outputs heads of tails.",
		"`-roll`: Randomly chooses a value between two numbers. Defaults to 0-100.",
		"`-lobby`: Creates a lobby with randomly chosen players.",
		"`-lobby_roles`: Creates a lobby with randomly chosen players and assigns each of them a role.",
	}
	return strings.Join(commands, "\n")
}

func roll_number() string {
	return strconv.Itoa(rand.Intn(101))
}
