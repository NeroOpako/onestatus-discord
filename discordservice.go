package main

import (
	"fmt"

	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
)

func setupDiscord() string {

	// Create a new session with GUILD_PRESENCES intent
	s, err := session.NewWithIntents("Bot "+secrets.DiscordAppToken, gateway.IntentGuilds|gateway.IntentGuildPresences)
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return "Error while initializing Discord Communication"
	}

	// Add a handler for presence updates
	s.AddHandler(func(e *gateway.PresenceUpdateEvent) {
		fmt.Printf("User %s updated presence: %s\n", e.User.ID, e.Activities)
		//sendPresence()
	})

	// Connect to Discord
	if err := s.Open(); err != nil {
		fmt.Println("Error opening connection:", err)
		return "Error while initializing Discord Communication"
	}
	defer s.Close()

	return ""
}
