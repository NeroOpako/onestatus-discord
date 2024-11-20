package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/diamondburned/arikawa/v3/discord"
)

var secrets Secrets

type Secrets struct {
	BSkyUserName     string `json:"bsky_user_name"`
	BSkyUserServer   string `json:"bsky_user_server"`
	BSkyUserPassword string `json:"bsky_user_password"`
	DiscordAppID     string `json:"discord_app_id"`
	DiscordAppToken  string `json:"discord_app_token"`
}

func loadSecrets() string {
	file, err := os.ReadFile("secrets.json")
	if err != nil {
		log.Fatalf("Error reading secrets file: %v", err)
		return "Error reading secrets file"
	}

	err = json.Unmarshal(file, &secrets)
	if err != nil {
		log.Fatalf("Error parsing secrets: %v", err)
		return "Error parsing secrets"
	}
	return ""
}

func main() {

	fmt.Println("Initializing service (1/3)")

	if err := loadSecrets(); err != "" {
		log.Fatal("Error while loading secrets.json, check if the file is correct.")
		return
	}

	fmt.Println("Initializing service (2/3)")

	if err := setupBlueSky(); err != "" {
		log.Fatal("Error while logging into BlueSky, check if your credentials are correct.")
		return
	}

	fmt.Println("Initializing service (3/3)")

	if err := setupDiscord(); err != "" {
		log.Fatal("Error while logging into Discord, check if your credentials are correct.")
		return
	}

	fmt.Println("Service is now running. Press CTRL+C to exit.")

	// Wait for a termination signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down service.")

}

func sendPresence(presence discord.Status) {
	createPost(skyUser, presence)
}
