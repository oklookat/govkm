package schema

import "net/url"

type (
	Tag struct {
		Name  string `json:"name"`
		Cover Image  `json:"cover"`
		APIID ID     `json:"apiId"`
	}

	Lyrics struct {
		License   string           `json:"license"`
		Intervals []LyricsInterval `json:"intervals"`
	}

	LyricsInterval struct {
		Begin     int    `json:"begin"`
		Countdown int    `json:"countdown,omitempty"`
		End       int    `json:"end"`
		Line      string `json:"line,omitempty"`
	}

	SocialLink struct {
		Avatar     Image  `json:"avatar"`
		SocialType string `json:"socialType"`
		Name       string `json:"name"`
		URL        string `json:"url"`
	}
)

type Image struct {
	URL         string `json:"url"`
	AccentColor string `json:"accentColor"`
	AvgColor    string `json:"avgColor"`
}

// Возвращает URL или nil, если спарсить не удалось.
func (e Image) GetUrl() *url.URL {
	parsed, err := url.Parse(e.URL)
	if err != nil {
		return nil
	}
	return parsed
}

type BlockOrder string

const (
	BlockOrderTracks     BlockOrder = "tracks"
	BlockOrderAlbums     BlockOrder = "albums"
	BlockOrderArtists    BlockOrder = "artists"
	BlockOrderPlaylists  BlockOrder = "playlists"
	BlockOrderPodcasts   BlockOrder = "podcasts"
	BlockOrderAudiobooks BlockOrder = "audiobooks"
	BlockOrderRadios     BlockOrder = "radios"
)
