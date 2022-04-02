package utils

import (
	"github.com/raitonoberu/ytmusic"
)

func GetSongData(songName string, index int) string {
	s := ytmusic.Search(songName)
	result, err := s.Next()
	PanicErr(err)
	if len(result.Tracks) == 0 {
		if len(result.Videos) == 0 {
			return ""
		}
		return result.Videos[index].VideoID
	}
	return result.Tracks[index].VideoID

}
