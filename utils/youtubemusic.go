package utils

import (
	"github.com/raitonoberu/ytmusic"
)

func GetSongData(songName string) string {
	s := ytmusic.Search(songName)
	result, err := s.Next()
	PanicErr(err)
	return result.Tracks[0].VideoID
}
