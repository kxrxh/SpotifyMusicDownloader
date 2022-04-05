package core

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	PlayListID   string
	ClientID     string
	ClientSecret string
	FolderName   string
	AppConfig    Config
)

type Config struct {
	Multithreading     bool
	CustomClientID     string
	CustomClientSecret string
}

func Init(playlistIdByUser string) {
	PlayListID = playlistIdByUser

	ClientSecret = "YOUR_SECRET"
	ClientID = "YOUR_ID"
	if ClientID == "" {
		log.Fatal("Unable to find client id.")
	}
	if ClientSecret == "" {
		log.Fatal("Unable to find client secret.")
	}
	// Parsing config file.
	_, err := toml.DecodeFile("app.toml", &AppConfig)
	if err != nil {
		AppConfig = Config{Multithreading: false, CustomClientID: "", CustomClientSecret: ""}
		log.Printf("Unable to decode config file: %v.\nMultithreading is disabled!\n", err)
	}
	// Setting up custom ClientID and ClientSecret (for devs).
	if AppConfig.CustomClientSecret != "" && AppConfig.CustomClientID != "" {
		ClientSecret = AppConfig.CustomClientSecret
		ClientID = AppConfig.CustomClientID
		log.Println("Using custom ClientID and ClientSecret!")
	}
}
