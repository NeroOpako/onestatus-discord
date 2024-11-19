package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

var accessToken string = ""
var secrets Secrets

type Secrets struct {
	BSKYTestUsername string `json:"bsky_test_username"`
	BSKYTestPassword string `json:"bsky_test_password"`
}

type LoginResponse struct {
	AccessJwt string `json:"accessJwt"`
}

type LoginPayload struct {
	Identifier string
	Password   string
}

func setup() string {
	if secrets.BSKYTestUsername == "" {
		file, err := os.ReadFile("secrets.json")
		if err != nil {
			log.Fatalf("Error reading secrets file: %v", err)
		}
		err = json.Unmarshal(file, &secrets)
		if err != nil {
			log.Fatalf("Error parsing secrets: %v", err)
		}
	}
	if accessToken == "" {
		bskySetup(LoginPayload{Identifier: secrets.BSKYTestUsername, Password: secrets.BSKYTestPassword})
	}
	return fetchProfile()
}

func bskySetup(userPayload LoginPayload) {
	client := resty.New()

	// Login credentials
	payload := map[string]string{
		"identifier": userPayload.Identifier, // Replace with your handle
		"password":   userPayload.Password,   // Replace with your password
	}

	// Make login request
	resp, err := client.R().
		SetBody(payload).
		SetResult(&LoginResponse{}).
		Post("https://bsky.social/xrpc/com.atproto.server.createSession")

	if err != nil {
		log.Fatal("Error:", err)
	}

	result := resp.Result().(*LoginResponse)
	fmt.Println("Access Token:", result.AccessJwt)
	accessToken = result.AccessJwt
}

func fetchProfile() string {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+accessToken).
		Get("https://bsky.social/xrpc/app.bsky.actor.getProfile?actor=neropako.dev")

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Println("Response:", resp.String())
	return resp.String()
}
