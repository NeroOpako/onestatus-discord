package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hugolgst/rich-go/client"
)

var secrets Secrets

type Secrets struct {
	DiscordAppID string `json:"discord_app_id"`
}

type SkyStatusDiscordSetupResponse struct {
	Presence string
	ErrorMsg string
}

func loadSecrets() SkyStatusDiscordSetupResponse {
	file, err := os.ReadFile("secrets.json")
	if err != nil {
		log.Fatalf("Error reading secrets file: %v", err)
		return SkyStatusDiscordSetupResponse{ErrorMsg: "Error while initializing Discord Communication"}
	}

	var secrets Secrets

	err = json.Unmarshal(file, &secrets)
	if err != nil {
		log.Fatalf("Error parsing secrets: %v", err)
		return SkyStatusDiscordSetupResponse{ErrorMsg: "Error while initializing Discord Communication"}
	}
	return SkyStatusDiscordSetupResponse{ErrorMsg: ""}
}

func getPresence() SkyStatusDiscordSetupResponse {
	if response := loadSecrets(); response.ErrorMsg != "" {
		return response
	}

	if err := client.Login(secrets.DiscordAppID); err != nil {
		log.Fatalf("Failed to initialize Discord RPC: %v", err)
		return SkyStatusDiscordSetupResponse{ErrorMsg: "Error while initializing Discord Communication"}
	}

	// Query the current Rich Presence
	/* presence := client.getPresence

	log.Printf("Current Rich Presence: %+v\n", presence) */
	return SkyStatusDiscordSetupResponse{Presence: ""}
}
