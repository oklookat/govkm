package schema

import "strings"

type (
	TracksResponse struct {
		Tracks []Track `json:"tracks"`
	}

	TrackResponse struct {
		Track *Track `json:"track"`
	}
)

type Track struct {
	// (?) Трек недоступен / изъят / etc.
	IsRestricted bool `json:"isRestricted"`

	Album SimpleAlbum `json:"album"`

	// Трек с ненормативной лексикой?
	IsExplicit bool `json:"isExplicit"`

	// ID трека в ВК.
	VkAudioID ID `json:"vkAudioId"`

	// Можно ли включить радио по треку?
	IsRadioCapable bool `json:"isRadioCapable"`

	Lyrics *Lyrics `json:"lyrics"`

	File string `json:"file"`

	// Может быть пустым, даже если IsLegal == true.
	Artist SimpleArtist `json:"artist"`

	IsAdded bool `json:"isAdded"`

	// Имя артиста.
	ArtistDisplayName string `json:"artistDisplayName"`

	// ID трека.
	APIID ID `json:"apiId"`

	// Трек загружен официально?
	IsLegal bool `json:"isLegal"`

	Cover Image `json:"cover"`

	// Длительность в секундах.
	Duration int `json:"duration"`

	// Лайкнут?
	IsLiked bool `json:"isLiked"`

	// (?) Хеш для того чтобы поделиться треком.
	ShareHash string `json:"shareHash"`

	// Название трека.
	Name string `json:"name"`

	// Размер в байтах.
	Size int `json:"size"`

	Counts *struct {
		Play int `json:"play"`
	} `json:"counts"`

	Hls string `json:"hls"`

	Artists []SimpleArtist `json:"artists"`

	Permissions struct {
		Reason string `json:"reason"`
		Permit bool   `json:"permit"`
	} `json:"permissions"`
}

// Обложка трека сгенерирована?
// Трек может быть официально загружен, но быть без обложки.
func (t Track) IsDefaultCover() bool {
	return len(t.Cover.URL) == 0 || strings.Contains(t.Cover.URL, "default_cover")
}
