package utils

import (
	"fmt"
	"github.com/KXRXH/SpotifyMusicDownloader/core"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
	"os/exec"
)

func DownloadSong(songUrl string) {
	client := youtube.Client{}
	video, err := client.GetVideo(songUrl)
	PanicErr(err)
	format := video.Formats.WithAudioChannels()
	if len(format) == 0 {
		fmt.Println("Unable to download:", songUrl)
		return
	}
	PanicErr(err)
	stream, _, err := client.GetStream(video, &format[0])
	PanicErr(err)
	filePath := video.Author[:len(video.Author)-5] + video.Title
	file, err := os.Create("./tmp/" + filePath + ".mp4")
	PanicErr(err)
	defer file.Close()
	_, err = io.Copy(file, stream)
	PanicErr(err)
	cmd := exec.Command("ffmpeg.exe", "-i", "./tmp/"+filePath+".mp4", core.FolderName+filePath+".mp3")
	FatalErr(cmd.Run())
}
