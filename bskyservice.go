package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

var skyUser BSkyLoggedUser

type BSkyLoggedUser struct {
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
	skyUser = BSkyLoggedUser{Server: secrets.BSkyUserServer, UserName: secrets.BSkyUserName}
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

func updateStatus(record Record) string {
	// Initialize Resty client
	client := resty.New()

	// Prepare the request payload
	payload := map[string]interface{}{
		"collection": "dev.neropako.onestatus.status", // Collection to save the record
		"repo":       skyUser.Did,                     // Replace with the user's DID or handle
		"record":     record,                          // The actual record data
	}

	// Make the API call
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+skyUser.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(skyUser.Server + "/xrpc/com.atproto.repo.createRecord")

	if err != nil {
		log.Fatalf("Error creating record: %v", err)
		return "Error while sending update"
	}

	// Check the response
	if resp.IsError() {
		log.Fatalf("API call failed: %s", resp.String())
		return "Error while sending update"
	}

	// Print the result
	fmt.Printf("Record created successfully: %s\n", resp.String())

	return ""
}
