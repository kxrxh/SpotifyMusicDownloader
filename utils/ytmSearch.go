package utils

import (
	"github.com/raitonoberu/ytmusic"
)

func GetSongData(songName string) string {
	s := ytmusic.Search(songName)
	result, err := s.Next()
	PanicErr(err)
	track := result.Tracks[0]
	return track.VideoID

}
