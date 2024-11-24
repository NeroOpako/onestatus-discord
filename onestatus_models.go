package main

type Record struct {
	CreatedAt    string `json:"createdAt"`
	App          string `json:"app"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	ActivityType string `json:"activityType"`
}

// ActivityType
const (
	DISCORD_PLAYING   = "DISCORD_PLAYING"
	DISCORD_LISTENING = "DISCORD_LISTENING"
	DISCORD_WATCHING  = "DISCORD_WATCHING"
	DISCORD_STREAMING = "DISCORD_STREAMING"
	DISCORD_CUSTOM    = "DISCORD_CUSTOM"
)
