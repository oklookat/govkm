package schema

type (
	ArtistsResponse struct {
		Artists []Artist `json:"artists"`
	}

	SimpleArtistsResponse struct {
		Artists []SimpleArtist `json:"artists"`
	}

	ArtistResponse struct {
		Artist *Artist `json:"artist"`
	}

	SimpleArtist struct {
		// Аватар артиста сгенерирован?
		IsAutoGenCover bool `json:"isAutoGenCover"`

		// ID артиста.
		APIID ID `json:"apiId"`

		// Аватар.
		Avatar Image `json:"avatar"`

		// Имя.
		Name string `json:"name"`
	}

	Artist struct {
		SimpleArtist

		// (?) Жанровые теги.
		Tags []Tag `json:"tags"`

		ShareHash string `json:"shareHash"`

		AddedAt UnixTime `json:"addedAt"`

		IsLiked bool `json:"isLiked"`

		HasTracks bool `json:"hasTracks"`

		IsRadioCapable bool `json:"isRadioCapable"`

		SocialLinks []SocialLink `json:"socialLinks"`

		Counts struct {
			NewTrack int `json:"newTrack"`
			Like     int `json:"like"`
			Play     int `json:"play"`
			Album    int `json:"album"`
		} `json:"counts"`

		UmaTags string `json:"umaTags"`

		UpdatedAt UnixTime `json:"updatedAt"`

		LastAlbumID ID `json:"lastAlbumId"`

		Bio interface{} `json:"bio"`
	}
)
