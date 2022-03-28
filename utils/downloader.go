package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadSong(fileName string, songUrl string) (err error) {

	out, err := os.Create(fileName + ".mp3")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(songUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
