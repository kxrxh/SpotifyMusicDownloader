package spotify

import (
	"context"
	"fmt"
	"github.com/KXRXH/SpotifyMusicDownloader/core"
	"github.com/KXRXH/SpotifyMusicDownloader/utils"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

func ParseSpotifyPlayList() []string {

	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     core.ClientID,
		ClientSecret: core.ClientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Couldn't get token: %v", err)
	}
	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	var (
		songs []string
		song  string
	)
	playlistInfo, err := client.GetPlaylist(ctx, spotify.ID(core.PlayListID))
	if err != nil {
		log.Fatalln(err)
	}

	// Getting track list of the spotify playlist.
	playListTracks, err := client.GetPlaylistTracks(ctx, spotify.ID(core.PlayListID))
	if err != nil {
		log.Fatalln(err)
	}

	// Create a playlist folder with name of the playlist (only if we do not have folder name from .env).
	if core.FolderName == "" {
		core.FolderName = "./" + playlistInfo.Name + "/"
	}

	// Getting songs names.
	for i := 0; i < playListTracks.Total; i++ {
		song = fmt.Sprintf("%s - %s", playListTracks.Tracks[i].Track.Artists[0].Name,
			playListTracks.Tracks[i].Track.Name)
		songs = append(songs, song)
	}
	fmt.Printf("Starting downloading songs from %v%s%v%v playlist.\nTotal tracks: %v%d%v\n", utils.Red, playlistInfo.Name,
		utils.Reset, utils.Green, utils.Red, playlistInfo.Tracks.Total, utils.Reset)
	return songs
}
