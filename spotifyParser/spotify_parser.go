package spotifyParser

import (
	"context"
	"fmt"
	"github.com/KXRXH/music/core"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

func ParseSpotifyPlayList(playlistId string) []string {

	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     core.ClientID,
		ClientSecret: core.ClientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	var (
		songs []string
		song  string
	)
	// handle playlist results
	playListData, _ := client.GetPlaylistTracks(ctx, spotify.ID(playlistId))
	for i := 0; i < playListData.Total; i++ {
		song = fmt.Sprintf("%s - %s", playListData.Tracks[i].Track.Artists[0].Name,
			playListData.Tracks[i].Track.Name)
		songs = append(songs, song)
	}
	return songs
}
