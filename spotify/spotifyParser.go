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

type Track struct {
	Artist string
	Name   string
}

func ParseSpotifyPlayList() []Track {

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
	playlistInfo, err := client.GetPlaylist(ctx, spotify.ID(core.PlayListID))
	if err != nil {
		log.Fatalln(err)
	}

	// Create a playlist folder with name of the playlist (only if we do not have folder name from .env).
	if core.FolderName == "" {
		core.FolderName = "./" + playlistInfo.Name + "/"
	}
	var tracks []Track
	getTrackList(&ctx, *client, &tracks, playlistInfo.Tracks.Total)
	fmt.Printf("Starting downloading songs from %v%s%v%v playlist.\nTotal tracks: %v%d%v\n", utils.Red,
		playlistInfo.Name, utils.Reset, utils.Green, utils.Red, playlistInfo.Tracks.Total, utils.Reset)
	return tracks
}

func getTrackList(ctx *context.Context, client spotify.Client, trackList *[]Track, numOfTracks int) {
	// Getting track list of the spotify playlist.
	var song Track
	c := 0
	for numOfTracks > 0 {
		playListTracks, err := client.GetPlaylistTracks(*ctx, spotify.ID(core.PlayListID), spotify.Offset(c*100))
		if err != nil {
			log.Fatalln(err)
		}
		for _, item := range playListTracks.Tracks {
			song = Track{
				Artist: item.Track.Artists[0].Name,
				Name:   item.Track.Name,
			}
			*trackList = append(*trackList, song)
		}
		c++
		numOfTracks /= 100
	}
}
