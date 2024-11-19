package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

var accessToken string = ""

type LoginResponse struct {
	AccessJwt string `json:"accessJwt"`
}

type LoginPayload struct {
	Identifier string
	Password   string
}

func setup() string {
	if accessToken == "" {
		bskySetup(LoginPayload{Identifier: "neropako.dev", Password: "xfbv-5op3-7pqy-emqh"})
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
