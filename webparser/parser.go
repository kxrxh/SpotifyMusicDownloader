package webparser

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

type SongData struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Url    string `json:"url"`
}

func ParseWeb(inputSong string, data *SongData) {
	inputData := strings.ToLower(strings.Replace(inputSong, " ", "+", 10000))
	c := colly.NewCollector()

	c.OnHTML("div.wrapper-tracklist.muslist", func(e *colly.HTMLElement) {
		scrapedData := e.ChildAttr("div.track.mustoggler", "data-musmeta")
		err := json.Unmarshal([]byte(scrapedData), &data)
		if err != nil {
			log.Fatalf("Unable to parse song data: %v", err)
		}
	})

	err := c.Visit("https://ruo.morsmusic.org/search/" + inputData)
	if err != nil {
		log.Fatal("Visiting error:", err)
	}

}
