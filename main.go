package main

import (
	"fmt"
	"github.com/KXRXH/SpotifyMusicDownloader/core"
	"github.com/KXRXH/SpotifyMusicDownloader/spotify"
	"github.com/KXRXH/SpotifyMusicDownloader/utils"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var (
		wg              sync.WaitGroup
		userInput       string
		downloadCounter int
	)
	fmt.Printf("%vPlaylist link or id%v: %v", utils.Red, utils.Reset, utils.Green)
	_, err := fmt.Scanf("%s", &userInput)
	utils.FatalErr(err)

	// Getting playlist id if userInput was url.
	if strings.Contains(userInput, "https://") {
		userInput = userInput[strings.LastIndex(userInput, "/")+1:]
	}
	// Initializing global variables.
	core.Init(userInput)

	// Getting []string with tracks.
	songsList := spotify.ParseSpotifyPlayList()

	// Creating playlist folder and temp folder for .mp4 files.
	utils.FatalErr(os.MkdirAll(core.FolderName, os.ModePerm))
	utils.FatalErr(os.MkdirAll("./tmp/", os.ModePerm))

	// Preparing data
	downloadCounter = 0
	startTime := time.Now().Unix()

	for _, item := range songsList {
		// Fixing names of user's tracks.
		trackFullName := item.Artist + " - " + item.Name
		songUrl := utils.GetSongData(trackFullName)
		if songUrl != "" {
			wg.Add(1)
			go startDownload(&wg, &downloadCounter, trackFullName, len(songsList), songUrl)
		}
	}
	wg.Wait()
	utils.PanicErr(os.RemoveAll("./tmp/"))
	fmt.Printf("%vSuccess! Finished in %d sec.%v", utils.Green, time.Now().Unix()-startTime, utils.Reset)
}

func startDownload(wg *sync.WaitGroup, n *int, item string, numOfTracks int, songUrl string) {
	defer wg.Done()
	fmt.Printf("Downloading: %v%s%v\n", utils.Blue, item, utils.Reset)
	utils.DownloadSong(songUrl)
	*n++
	fmt.Printf("%vDownloading is finnished%v [%v%d%v/%v%d%v]: %v%s%v\n", utils.Green, utils.Reset, utils.Green, *n,
		utils.Reset, utils.Yellow, numOfTracks, utils.Reset, utils.Blue, item, utils.Reset)
}
