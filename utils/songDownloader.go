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

func DownloadSong(track string, songUrl string) bool {
	var filePath string
	client := youtube.Client{}

	video, err := client.GetVideo(songUrl)
	PanicErr(err)
	// Getting formats with audio channels.
	format := video.Formats.WithAudioChannels()
	// Retry
	if len(format) == 0 {
		GetSongData(track, 1)
		video, err = client.GetVideo(songUrl)
		if err != nil {
			log.Println("Unable to download:", track)
			return false
		}
		// Getting formats with audio channels.
		format = video.Formats.WithAudioChannels()
		if len(format) == 0 {
			log.Println("Unable to download:", track)
			return false
		}
	}
	stream, _, err := client.GetStream(video, &format[0])
	PanicErr(err)

	// Preparing file name.
	filePath = track
	filePath = strings.ReplaceAll(filePath, "/", "")
	filePath = strings.ReplaceAll(filePath, ":", " ")
	filePath = strings.ReplaceAll(filePath, "*", " ")

	// Creating mp4 file and writing to it
	file, err := os.Create("./tmp/" + filePath + ".mp4")
	PanicErr(err)
	defer func(file *os.File) {
		err = file.Close()
		PanicErr(err)
	}(file)
	_, err = io.Copy(file, stream)
	PanicErr(err)

	// Getting audio stream of mp4 file and creating mp3 file at playlist dir.
	cmd := exec.Command("ffmpeg.exe", "-i", "./tmp/"+filePath+".mp4", "-y", core.FolderName+filePath+".mp3")
	PanicErr(cmd.Run())
	return true
}
