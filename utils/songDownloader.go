package utils

import (
	"fmt"
	"github.com/KXRXH/SpotifyMusicDownloader/core"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
	"os/exec"
	"strings"
)

func DownloadSong(songUrl string) {
	var filePath string
	client := youtube.Client{}
	video, err := client.GetVideo(songUrl)
	PanicErr(err)

	// Getting formats with audio channels.
	format := video.Formats.WithAudioChannels()
	if len(format) == 0 {
		fmt.Println("Unable to download file:", songUrl)
		return
	}
	stream, _, err := client.GetStream(video, &format[0])
	PanicErr(err)

	// Editing song name (Deleting "Topic" in author`s name) (video.Author can be equal "Author Name - Topic").
	// Preparing file name.
	if strings.Contains(video.Author, "Topic") {
		filePath = video.Author[:len(video.Author)-5] + video.Title
	} else {
		filePath = video.Author + video.Title
	}

	// Creating mp4 file and writing to it
	file, err := os.Create("./tmp/" + filePath + ".mp4")
	defer func(file *os.File) {
		err = file.Close()
		PanicErr(err)
	}(file)
	_, err = io.Copy(file, stream)
	PanicErr(err)

	// Getting audio stream of mp4 file and creating mp3 file at playlist dir.
	cmd := exec.Command("ffmpeg.exe", "-i", "./tmp/"+filePath+".mp4", "-y", core.FolderName+filePath+".mp3")
	FatalErr(cmd.Run())
}
