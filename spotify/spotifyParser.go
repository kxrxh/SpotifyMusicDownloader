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

func GetTracksFromSpotifyPlaylist() []Track {

	// Client initialization
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     core.ClientID,
		ClientSecret: core.ClientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Could not get token for client: %v\n", err)
	}
	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	// This variable contains the name of the playlist and other information about the playlist.
	playlistInfo, err := client.GetPlaylist(ctx, spotify.ID(core.PlayListID))
	if err != nil {
		log.Fatalf("\nCould not get information about playlist with id: %s\n", core.PlayListID)
	}

	// Create a playlist folder with name of the playlist (only if we do not have folder name from .env).
	if core.FolderName == "" {
		core.FolderName = "./" + playlistInfo.Name + "/"
	}

	var tracks []Track

	getTrackList(&ctx, *client, &tracks, playlistInfo.Tracks.Total)
	fmt.Printf("Starting downloading songs from %v%s%v%v.\nTotal tracks number: %v%d%v\n", utils.Red,
		playlistInfo.Name, utils.Reset, utils.Green, utils.Red, playlistInfo.Tracks.Total, utils.Reset)
	return tracks
}

// getTrackList adds tracks data from the playlist to an array. It edits trackList pointer.
func getTrackList(ctx *context.Context, client spotify.Client, trackList *[]Track, numOfTracks int) {
	var trackItem Track
	// Because the maximum number of track which we can get from a 1 request is 100
	// We need to download playlist by parts. (1 part is <= 100 tracks).
	c := 0
	for numOfTracks > 0 {
		playListTracks, err := client.GetPlaylistTracks(*ctx, spotify.ID(core.PlayListID), spotify.Offset(c*100))
		if err != nil {
			log.Fatalf("Could not get tracks from the playlist: %s\n", core.PlayListID)
		}
		for _, item := range playListTracks.Tracks {
			trackItem = Track{
				Artist: item.Track.Artists[0].Name,
				Name:   item.Track.Name,
			}
			*trackList = append(*trackList, trackItem)
		}
		c++
		numOfTracks -= 100
	}
}
