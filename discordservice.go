package main

import (
	"fmt"

	"github.com/diamondburned/arikawa/discord"
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
		if len(e.Activities) > 0 {
			action := ""
			switch e.Activities[0].Type {
			case discord.GameActivity:
				action = "playing"
			case discord.WatchingActivity:
				action = "watching"
			case discord.StreamingActivity: // Multiple cases for the same action
				action = "streaming"
			case discord.ListeningActivity: // Multiple cases for the same action
				action = "listening"
			case discord.CustomActivity: // Multiple cases for the same action
				action = ""
			default: // Optional default case
				action = ""
			}
			fmt.Printf("User is %s %s\n", action, e.Activities[0].Name)
			//sendPresence()
		}
	})

	// Connect to Discord
	if err := s.Open(); err != nil {
		fmt.Println("Error opening connection:", err)
		return "Error while initializing Discord Communication"
	}

	return ""
}
