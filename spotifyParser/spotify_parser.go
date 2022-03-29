package spotifyParser

import (
	"context"
	"fmt"
	"github.com/KXRXH/SpotifyMusicDownloader/core"
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
	playListTracks, _ := client.GetPlaylistTracks(ctx, spotify.ID(playlistId))
	if core.FolderName == "" {
		playList, _ := client.GetPlaylist(ctx, spotify.ID(playlistId))
		core.FolderName = "./" + playList.Name + "/"
	}
	for i := 0; i < playListTracks.Total; i++ {
		song = fmt.Sprintf("%s - %s", playListTracks.Tracks[i].Track.Artists[0].Name,
			playListTracks.Tracks[i].Track.Name)
		songs = append(songs, song)
	}
	return songs
}
