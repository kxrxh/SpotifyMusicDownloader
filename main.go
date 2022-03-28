package main

import (
	"fmt"
	"github.com/KXRXH/music/core"
	"github.com/KXRXH/music/spotifyParser"
	"github.com/KXRXH/music/utils"
	"github.com/KXRXH/music/webparser"
	"log"
	"os"
	"time"
)

func main() {
	var (
		data      webparser.SongData
		inputData string
	)
	core.Init()
	songsList := spotifyParser.ParseSpotifyPlayList(core.PlayListID)
	fmt.Printf("Starting downloading songs from the playlist[https://open.spotify.com/playlist/%v]\n",
		core.PlayListID)
	startTime := time.Now().Unix()
	for n, item := range songsList {
		inputData = item
		webparser.ParseWeb(inputData, &data)

		downloadUrl := "https://ruo.morsmusic.org" + data.Url
		fmt.Printf("Downloading [%v%d%v/%v%d%v]: %v%s%v\n", utils.Green, n+1,
			utils.Reset, utils.Yellow, len(songsList), utils.Reset, utils.Blue, item, utils.Reset)

		err := os.MkdirAll(core.FolderName, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		err = utils.DownloadSong(core.FolderName+item, downloadUrl)
		if err != nil {
			log.Println("Unable to download file:", err)
		}
	}
	fmt.Printf("%vSuccess! Finished in %d sec.%v", utils.Green, time.Now().Unix()-startTime, utils.Reset)
}
