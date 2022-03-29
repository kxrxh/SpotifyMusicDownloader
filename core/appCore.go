package core

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	PlayListID   string
	ClientID     string
	ClientSecret string
	FolderName   string
)

func Init(playlistIdByUser string) {

	// Reading variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PlayListID = playlistIdByUser
	ClientID = os.Getenv("CLIENT_ID")
	if ClientID == "" {
		log.Fatal("Unable to find client id.")
	}
	ClientSecret = os.Getenv("CLIENT_SECRET")
	if ClientSecret == "" {
		log.Fatal("Unable to find client secret.")
	}
	FolderName = os.Getenv("FOLDER_NAME")
}
