package utils

import (
	"github.com/KXRXH/SpotifyMusicDownloader/core"
	"github.com/kkdai/youtube/v2"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// DownloadTrack will download a track by its URL.
// It creates mp4 with track name then, gets audio stream and creates an mp3 file with that stream.
// Returns an error if the error is fatal.
func DownloadTrack(track string, songUrl string) error {
	var filePath string
	client := youtube.Client{}

	video, err := client.GetVideo(songUrl)
	if err != nil {
		log.Printf("Unable to download: %s\nError: %v\n", track, err)
		return err
	}
	// Getting formats with audio channels.
	format := video.Formats.WithAudioChannels()
	// Retry if unable to download track
	if len(format) == 0 {
		// Getting next track from search.
		songUrl = GetSongData(track, 1)
		video, err = client.GetVideo(songUrl)
		if err != nil {
			log.Printf("Unable to download: %s\nError: %v\n", track, err)
			return err
		}
		format = video.Formats.WithAudioChannels()
		if len(format) == 0 {
			log.Printf("Unable to download: %s\nError: %v\n", track, err)
			return err
		}
	}
	stream, _, err := client.GetStream(video, &format[0])
	if err != nil {
		log.Printf("Unable to download: %s\nError: %v\n", track, err)
		return err
	}

	// Refactoring track name
	filePath = track
	filePath = strings.ReplaceAll(filePath, "/", "_")
	filePath = strings.ReplaceAll(filePath, ":", "_")
	filePath = strings.ReplaceAll(filePath, "*", "_")
	filePath = strings.ReplaceAll(filePath, "?", "_")

	// Creating mp4 file and writing to it
	file, err := os.Create("./tmp/" + filePath + ".mp4")
	if err != nil {
		log.Printf("Unable to create mp4 file: %s\nError: %v\n", filePath, err)
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Printf("Unable to close mp4 file: %s\nError: %v\n", filePath, err)
		}
	}(file)
	_, err = io.Copy(file, stream)
	if err != nil {
		log.Printf("Copying stream to file error: %s\nError: %v\n", filePath, err)
		return err
	}

	// Getting audio stream of mp4 file and creating mp3 file at playlist dir.
	cmd := exec.Command("ffmpeg.exe", "-i", "./tmp/"+filePath+".mp4", "-y", core.FolderName+filePath+".mp3")
	err = cmd.Run()
	if err != nil {
		log.Printf("Unable to create mp3 file with audio stream: %s\nError: %v\n", filePath, err)
		return err
	}
	return nil
}
