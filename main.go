package main

import (
	"fmt"
	"github.com/KXRXH/SpotifyMusicDownloader/core"
	"github.com/KXRXH/SpotifyMusicDownloader/spotifyParser"
	"github.com/KXRXH/SpotifyMusicDownloader/utils"
	"os"
	"sync"
	"time"
)

func main() {
	var (
		wg      sync.WaitGroup
		inputId string
	)
	fmt.Printf("Playlist id: ")
	_, err := fmt.Scanf("%s", &inputId)
	utils.FatalErr(err)
	core.Init(inputId)
	songsList := spotifyParser.ParseSpotifyPlayList(core.PlayListID)
	fmt.Printf("Starting downloading songs from the playlist[https://open.spotify.com/playlist/%v]\n",
		core.PlayListID)
	startTime := time.Now().Unix()
	utils.FatalErr(os.MkdirAll(core.FolderName, os.ModePerm))
	utils.FatalErr(os.MkdirAll("./tmp/", os.ModePerm))
	counter := 0
	for _, item := range songsList {
		songUrl := utils.GetSongData(item)
		wg.Add(1)
		go startDownload(&wg, &counter, item, songsList, songUrl)

	}
	wg.Wait()
	utils.PanicErr(os.RemoveAll("./tmp/"))
	fmt.Printf("%vSuccess! Finished in %d sec.%v", utils.Green, time.Now().Unix()-startTime, utils.Reset)
}

func startDownload(wg *sync.WaitGroup, n *int, item string, songsList []string, songUrl string) {
	defer wg.Done()
	fmt.Printf("Downloading: %v%s%v\n", utils.Blue, item, utils.Reset)
	utils.DownloadSong(songUrl)
	*n++
	fmt.Printf("%vDownloading is finnished%v [%v%d%v/%v%d%v]: %v%s%v\n", utils.Green, utils.Reset, utils.Green, *n,
		utils.Reset, utils.Yellow, len(songsList), utils.Reset, utils.Blue, item, utils.Reset)
}
