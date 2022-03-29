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

func Init(input string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PlayListID = input
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
