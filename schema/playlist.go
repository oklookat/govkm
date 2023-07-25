package schema

type (
	PlaylistResponse struct {
		Playlist *Playlist `json:"playlist"`
	}

	PlaylistsResponse struct {
		Playlists []Playlist `json:"playlists"`
	}

	Playlist struct {
		// Плейлист с лайкнутыми треками?
		IsFavorite bool `json:"isFavorite"`

		UpdatedAt UnixTime `json:"updatedAt"`

		SpecialCover Image `json:"specialCover"`

		Name string `json:"name"`

		Owner SimpleUser `json:"owner"`

		Duration int `json:"duration"`

		IsEnhanced bool `json:"isEnhanced"`

		IsRadioCapable bool `json:"isRadioCapable"`

		ShareHash string `json:"shareHash"`

		IsDownloads bool `json:"isDownloads"`

		IsOldBoom bool `json:"isOldBoom"`

		APIID ID `json:"apiId"`

		Counts struct {
			Like  int `json:"like"`
			Track int `json:"track"`
			Play  int `json:"play"`
		} `json:"counts"`

		IsExplicit bool `json:"isExplicit"`

		Celebrity bool `json:"celebrity"`

		IsLiked bool `json:"isLiked"`

		BadSync bool `json:"badSync"`

		Source PlaylistSource `json:"source"`

		Type PlaylistType `json:"type"`

		IsDefault bool `json:"isDefault"`

		AddedAt UnixTime `json:"addedAt"`

		Cover Image `json:"cover"`
	}
)

type PlaylistSource string

const (
	PlaylistSourceMoosic PlaylistSource = "moosic"
	PlaylistSourceVK     PlaylistSource = "vk"
)

type PlaylistType string

const (
	PlaylistTypeCommon PlaylistType = "common"
)
