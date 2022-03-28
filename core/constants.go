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

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PlayListID = os.Getenv("PLAYLIST_ID")
	if PlayListID == "" {
		log.Fatal("Unable to find playlist id.")
	}
	ClientID = os.Getenv("CLIENT_ID")
	if ClientID == "" {
		log.Fatal("Unable to find client id.")
	}
	ClientSecret = os.Getenv("CLIENT_SECRET")
	if ClientSecret == "" {
		log.Fatal("Unable to find client secret.")
	}
	FolderName = os.Getenv("FOLDER_NAME")
	if FolderName == "" {
		FolderName = "./songs/"
	}

}
