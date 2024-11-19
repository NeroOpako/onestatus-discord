package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Welcome to my simple terminal app!")

	fmt.Print("Please enter your server [Default: bsky.social]: ")

	// Read user input
	var server string
	_, err := fmt.Scanln(&server)
	if err != nil {
		log.Fatal(err)
		return
	}

	if server == "" {
		server = "bsky.social"
	} else {
		server = "https://" + server
	}

	fmt.Print("Please enter your username : ")

	// Read user input
	var username string
	_, err = fmt.Scanln(&username)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Print("Please enter your password [Use an app password!] : ")

	// Read user input
	var password string
	_, err = fmt.Scanln(&password)
	if err != nil {
		log.Fatal(err)
		return
	}

	skyStatusResponse := setup(SkyStatusSetupRequest{Server: server, Identifier: username, Password: password})
	if skyStatusResponse.ErrorMsg != "" {
		log.Fatal("Error while logging into BlueSky, check if your credentials are correct.")
		return
	}

	discordResponse := getPresence()
	if discordResponse.ErrorMsg != "" {
		log.Fatal("Error while reading Discord status")
		return
	}

	if discordResponse.Presence != "" {
		createPost(skyStatusResponse.LoggedUser, discordResponse.Presence)
	}

}
