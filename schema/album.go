package schema

type (
	AlbumsResponse struct {
		Albums []Album `json:"albums"`
	}

	AlbumResponse struct {
		Album *Album `json:"album"`
	}

	SimpleAlbum struct {
		Cover Image  `json:"cover"`
		Name  string `json:"name"`
		APIID ID     `json:"apiId"`
	}

	Album struct {
		SimpleAlbum

		Exclusive bool `json:"exclusive"`

		Year int `json:"year"`

		Artists []Artist `json:"artists"`

		Tags []Tag `json:"tags"`

		Counts struct {
			Like  int `json:"like"`
			Track int `json:"track"`
		} `json:"counts"`

		Permissions struct {
			Reason string `json:"reason"`
			Permit bool   `json:"permit"`
		} `json:"permissions"`

		ShareHash string `json:"shareHash"`

		IsLiked bool `json:"isLiked"`

		Description string `json:"description"`

		ArtistDisplayName string `json:"artistDisplayName"`

		UpdatedAt UnixTime `json:"updatedAt"`

		AddedAt UnixTime `json:"addedAt"`

		Duration int `json:"duration"`

		IsRadioCapable bool `json:"isRadioCapable"`

		Types []AlbumType `json:"types"`

		IsExplicit bool `json:"isExplicit"`

		ReleaseDateTimestamp UnixTime `json:"releaseDateTimestamp"`

		UmaTags string `json:"umaTags"`
	}
)

type AlbumType string

func (e AlbumType) String() string {
	return string(e)
}

const (
	AlbumTypeSingle      AlbumType = "single"
	AlbumTypeAlbum       AlbumType = "album"
	AlbumTypeEp          AlbumType = "ep"
	AlbumTypeRemix       AlbumType = "remix"
	AlbumTypeCompilation AlbumType = "compilation"
)
