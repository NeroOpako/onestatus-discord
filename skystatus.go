package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

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

type SkyStatusSetupRequest struct {
	Identifier string
	Password   string
	Server     string
}

func setup(payload SkyStatusSetupRequest) SkyStatusSetupResponse {
	var response SkyStatusSetupResponse
	response.LoggedUser = SkyStatusLoggedUser{Server: payload.Server, UserName: payload.Identifier}
	if response = getAccessToken(payload, response); response.ErrorMsg != "" {
		return response
	}
	if response = getDid(response); response.ErrorMsg != "" {
		return response
	}
	return response
}

func getAccessToken(userPayload SkyStatusSetupRequest, response SkyStatusSetupResponse) SkyStatusSetupResponse {
	client := resty.New()

	// Login credentials
	payload := map[string]string{
		"identifier": userPayload.Identifier, // Replace with your handle
		"password":   userPayload.Password,   // Replace with your password
	}
	// Make login request
	resp, err := client.R().
		SetBody(payload).
		SetResult(&BSkyLoginResponse{}).
		Post(response.LoggedUser.Server + "/xrpc/com.atproto.server.createSession")

	if err != nil {
		log.Fatal("Error:", err)
		response.ErrorMsg = "Error while logging in"
		return response
	}
	result := resp.Result().(*BSkyLoginResponse)
	response.LoggedUser.AccessToken = result.AccessJwt
	return response
}

func getDid(response SkyStatusSetupResponse) SkyStatusSetupResponse {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+response.LoggedUser.AccessToken).
		SetResult(&BSkyProfileResponse{}).
		Get(response.LoggedUser.Server + "/xrpc/app.bsky.actor.getProfile?actor=" + response.LoggedUser.UserName)

	if err != nil {
		log.Fatal("Error:", err)
		response.ErrorMsg = "Error while retrieving user data"
		return response
	}

	result := resp.Result().(*BSkyProfileResponse)
	response.LoggedUser.Did = result.Did
	return response
}

func createPost(loggedUser SkyStatusLoggedUser, presence string) string {
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
