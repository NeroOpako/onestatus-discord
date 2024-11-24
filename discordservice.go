package main

import (
	"fmt"
	"time"

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
	s.AddHandler(getUpdatedPresence)

	// Connect to Discord
	if err := s.Open(); err != nil {
		fmt.Println("Error opening connection:", err)
		return "Error while initializing Discord Communication"
	}

	return ""
}

func getUpdatedPresence(e *gateway.PresenceUpdateEvent) {
	if len(e.Activities) > 0 {
		action := ""
		record := Record{CreatedAt: time.Now().UTC().Format(time.RFC3339)}
		switch e.Activities[0].Type {
		case discord.GameActivity:
			record.ActivityType = DISCORD_PLAYING
			record.App = "Discord"
			record.Title = e.Activities[0].Name
			record.Description = e.Activities[0].Details
			if e.Activities[0].Assets != nil {
				record.Image = e.Activities[0].Assets.LargeImage
			}
			action = "playing"
		case discord.WatchingActivity:
			record.ActivityType = DISCORD_WATCHING
			record.App = e.Activities[0].Name
			record.Title = e.Activities[0].Name
			record.Description = e.Activities[0].Details
			if e.Activities[0].Assets != nil {
				record.Image = e.Activities[0].Assets.LargeImage
			}
			action = "watching"
		case discord.StreamingActivity:
			record.ActivityType = DISCORD_STREAMING
			record.App = e.Activities[0].Name
			record.Title = e.Activities[0].Name
			record.Description = e.Activities[0].Details
			if e.Activities[0].Assets != nil {
				record.Image = e.Activities[0].Assets.LargeImage
			}
			action = "streaming"
		case discord.ListeningActivity:
			record.ActivityType = DISCORD_LISTENING
			record.App = e.Activities[0].Name
			record.Title = e.Activities[0].Name
			record.Description = e.Activities[0].Details
			if e.Activities[0].Assets != nil {
				record.Image = e.Activities[0].Assets.LargeImage
			}
			action = "listening"
		case discord.CustomActivity:
			record.ActivityType = DISCORD_CUSTOM
			record.App = e.Activities[0].Name
			record.Title = e.Activities[0].Name
			record.Description = e.Activities[0].Details
			if e.Activities[0].Assets != nil {
				record.Image = e.Activities[0].Assets.LargeImage
			}
			action = ""
		default:
			action = ""
		}
		fmt.Printf("User is %s %s\n", action, e.Activities[0].Name)
		updateStatus(record)
	}
}
