package main

import (
	"fmt"
	"log"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/go-resty/resty/v2"
)

var skyUser SkyStatusLoggedUser

type SkyStatusSetupResponse struct {
	LoggedUser SkyStatusLoggedUser
	ErrorMsg   string
}

type SkyStatusLoggedUser struct {
	Did         string
	AccessToken string
	UserName    string
	Server      string
}

type BSkyProfileResponse struct {
	Did string `json:"did"`
}

type BSkyLoginResponse struct {
	AccessJwt string `json:"accessJwt"`
}

func setupBlueSky() string {
	skyUser = SkyStatusLoggedUser{Server: secrets.BSkyUserServer, UserName: secrets.BSkyUserName}
	if response := getAccessToken(); response != "" {
		return response
	}
	if response := getDid(); response != "" {
		return response
	}
	return ""
}

func getAccessToken() string {
	client := resty.New()

	// Login credentials
	payload := map[string]string{
		"identifier": secrets.BSkyUserName,     // Replace with your handle
		"password":   secrets.BSkyUserPassword, // Replace with your password
	}
	// Make login request
	resp, err := client.R().
		SetBody(payload).
		SetResult(&BSkyLoginResponse{}).
		Post(skyUser.Server + "/xrpc/com.atproto.server.createSession")

	if err != nil {
		log.Fatal("Error:", err)
		return "Error while logging in"
	}
	result := resp.Result().(*BSkyLoginResponse)
	skyUser.AccessToken = result.AccessJwt
	return ""
}

func getDid() string {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+skyUser.AccessToken).
		SetResult(&BSkyProfileResponse{}).
		Get(skyUser.Server + "/xrpc/app.bsky.actor.getProfile?actor=" + skyUser.UserName)

	if err != nil {
		log.Fatal("Error:", err)
		return "Error while retrieving user data"
	}

	result := resp.Result().(*BSkyProfileResponse)
	skyUser.Did = result.Did
	return ""
}

func createPost(loggedUser SkyStatusLoggedUser, presence discord.Status) string {
	// Initialize Resty client
	client := resty.New()

	// Define the record data
	recordData := map[string]interface{}{
		"text":      presence,
		"createdAt": time.Now().UTC().Format(time.RFC3339), // Proper timestamp format
	}

	// Prepare the request payload
	payload := map[string]interface{}{
		"collection": "app.bsky.feed.post", // Collection to save the record
		"repo":       loggedUser.Did,       // Replace with the user's DID or handle
		"record":     recordData,           // The actual record data
	}

	// Make the API call
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+loggedUser.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(loggedUser.Server + "/xrpc/com.atproto.repo.createRecord")

	if err != nil {
		log.Fatalf("Error creating record: %v", err)
	}

	// Check the response
	if resp.IsError() {
		log.Fatalf("API call failed: %s", resp.String())
	}

	// Print the result
	fmt.Printf("Record created successfully: %s\n", resp.String())

	return resp.String()
}
