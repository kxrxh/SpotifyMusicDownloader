package main

import (
	"fmt"
	"github.com/KXRXH/SpotifyMusicDownloader/core"
	"github.com/KXRXH/SpotifyMusicDownloader/spotify"
	"github.com/KXRXH/SpotifyMusicDownloader/utils"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var (
		userInput    string
		trackCounter int
	)
	fmt.Printf("%vPlaylist link or id%v: %v", utils.Red, utils.Reset, utils.Green)
	_, err := fmt.Scanf("%s", &userInput)
	if err != nil {
		log.Fatalf("\nIncorrect user input: %v\n", err)
	}

	// Getting playlist id if userInput was url.
	if strings.Contains(userInput, "https://") {
		userInput = userInput[strings.LastIndex(userInput, "/")+1:]
	}
	// Initializing global variables.
	core.Init(userInput)

	// Getting array with tracks.
	trackList := spotify.GetTracksFromSpotifyPlaylist()

	// Creating playlist dir for .mp3 files and temp dir for .mp4 files.
	err = os.MkdirAll(core.FolderName, os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to create a folder with name: %s.\nError: %v\n", core.FolderName, err)
	}
	err = os.MkdirAll("./tmp/", os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to create a tmp directory.\nError: %v\n", err)
	}

	trackCounter = 0               // Counts number of downloaded tracks (without fatal errors).
	startTime := time.Now().Unix() // Start download time.

	// Check if multithreading is enabled or disabled.
	if core.AppConfig.Multithreading {
		// Download with multithreading (heavy load on the GPU), but faster.
		fmt.Printf("%vMultithreading is enabled!%v\n", utils.Blue, utils.Reset)
		var wg sync.WaitGroup
		for _, track := range trackList {
			fullNameOfTheTrack := track.Artist + " - " + track.Name
			songUrl := utils.GetSongData(fullNameOfTheTrack, 0)
			if songUrl != "" {
				wg.Add(1)
				go startDownloadWithMultithreading(&wg, &trackCounter, fullNameOfTheTrack, len(trackList), songUrl)
			}
		}
		wg.Wait()
	} else {
		fmt.Printf("%vMultithreading is disabled!%v\n", utils.Blue, utils.Reset)
		for _, track := range trackList {
			fullNameOfTheTrack := track.Artist + " - " + track.Name
			songUrl := utils.GetSongData(fullNameOfTheTrack, 0)
			if songUrl != "" {
				startDownload(&trackCounter, fullNameOfTheTrack, len(trackList), songUrl)
			}
		}
	}
	err = os.RemoveAll("./tmp/")
	if err != nil {
		log.Println("Unable to remove tmp directory.", err)
	}
	fmt.Printf("%vSuccess! Finished in %d sec.%v\n", utils.Green, time.Now().Unix()-startTime, utils.Reset)
	fmt.Printf("%vErrors: %d%v\n", utils.Red, len(trackList)-trackCounter, utils.Reset)
}

// startDownload starts common download. It is slow, but easier for your computer.
func startDownload(trackCounter *int, track string, totalNumberOfTracks int, songUrl string) {
	fmt.Printf("Downloading: %v%s%v\n", utils.Blue, track, utils.Reset)
	// If the track is downloaded then the counter is + 1
	err := utils.DownloadTrack(track, songUrl)
	if err != nil {
		fmt.Printf("%vDownloading is finnished with error: %s\nError: %v%v\n", utils.Red, track, err, utils.Reset)
		return
	}
	*trackCounter++
	fmt.Printf("%vDownloading is finnished%v [%v%d%v/%v%d%v]: %v%s%v\n", utils.Green, utils.Reset, utils.Green,
		*trackCounter, utils.Reset, utils.Yellow, totalNumberOfTracks, utils.Reset, utils.Blue, track, utils.Reset)

}

// startDownloadWithMultithreading starts download with multithreading. It faster than common download, but
// it causes heavy gpu load (with big playlist at least).
func startDownloadWithMultithreading(wg *sync.WaitGroup, trackCounter *int, track string, totalNumberOfTracks int,
	songUrl string) {
	defer wg.Done()
	fmt.Printf("Downloading: %v%s%v\n", utils.Blue, track, utils.Reset)
	// If the track is downloaded then the counter is + 1
	err := utils.DownloadTrack(track, songUrl)
	if err != nil {
		fmt.Printf("%vDownloading is finnished with error: %s\nError: %v%v\n", utils.Red, track, err, utils.Reset)
		return
	}
	*trackCounter++
	fmt.Printf("%vDownloading is finnished%v [%v%d%v/%v%d%v]: %v%s%v\n", utils.Green, utils.Reset, utils.Green,
		*trackCounter, utils.Reset, utils.Yellow, totalNumberOfTracks, utils.Reset, utils.Blue, track, utils.Reset)

}
