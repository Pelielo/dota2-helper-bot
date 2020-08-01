package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
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
	fmt.Printf("%v said: %v\n", s.State.User.Username, m.Content)

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if msg := m.Content; strings.HasPrefix(msg, Prefix) {
		fmt.Println(msg)
		fmt.Println(strings.Split(msg[len(Prefix):], " "))
		switch commands := strings.Split(msg[len(Prefix):], " "); {
		// Show documentation on lobby command
		case commands[0] == "lobby" && commands[1] == "help":
			s.ChannelMessageSend(m.ChannelID, "Help!")
		// Randomizes a lobby of 10, 11 or 12 people
		case commands[0] == "lobby":
			players := shuffle_array(commands[1:])
			s.ChannelMessageSend(m.ChannelID, strings.Join(players, "\n"))
		}
	}
}

func shuffle_array(a []string) []string {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}
