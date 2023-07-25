package schema

type (
	SearchResult struct {
		BlocksOrder []BlockOrder `json:"blocksOrder"`

		Playlists []Playlist `json:"playlists"`

		Podcasts []interface{} `json:"podcasts"`

		Radios []interface{} `json:"radios"`

		Afters struct {
			Radio interface{} `json:"radio"`
		} `json:"afters"`

		Artists []Artist `json:"artists"`

		Audiobooks []interface{} `json:"audiobooks"`

		Counts struct {
			Album  int `json:"album"`
			Artist int `json:"artist"`
			Track  int `json:"track"`
		} `json:"counts"`

		Tracks []Track `json:"tracks"`

		QueryIds struct {
			Track    string `json:"track"`
			Playlist string `json:"playlist"`
			Album    string `json:"album"`
			Artist   string `json:"artist"`
		} `json:"queryIds"`

		Albums []Album `json:"albums"`
	}

	SearchTrackResult struct {
		Tracks []Track `json:"tracks"`

		QueryIds struct {
			Track string `json:"track"`
		} `json:"queryIds"`

		Counts struct {
			Track int `json:"track"`
		} `json:"counts"`
	}

	SearchSuggestion struct {
		Suggestions []string `json:"suggestions"`

		Counts struct {
			Suggestion int `json:"suggestion"`
		} `json:"counts"`

		ObjectSuggestions []interface{} `json:"objectSuggestions"`
	}
)
