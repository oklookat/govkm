package schema

type (
	UserResponse struct {
		User User `json:"user"`
	}

	UserSettingsResponse struct {
		User UserSettings `json:"user"`
	}

	SimpleUser struct {
		Avatar Image `json:"avatar"`

		FirstName string `json:"firstName"`

		LastName string `json:"lastName"`

		IsOnline bool `json:"isOnline"`

		ShareHash string `json:"shareHash"`

		APIID ID `json:"apiId"`
	}

	User struct {
		SimpleUser

		Counts UserCounts `json:"counts"`

		Billing UserBilling `json:"billing"`

		ShowRecommendationsOnboardingInPreferences bool `json:"showRecommendationsOnboardingInPreferences"`

		OauthSource string `json:"oauthSource"`

		Cover Image `json:"cover"`

		PlaylistSyncProgress UserPlaylistSyncProgress `json:"playlistSyncProgress"`

		OauthID string `json:"oauthId"`

		UpdateTime UserUpdateTime `json:"updateTime"`

		OnLoginShowRecommendationsOnboarding bool `json:"onLoginShowRecommendationsOnboarding"`

		UserGeo string `json:"userGeo"`

		Tags []Tag `json:"tags"`

		HasFeed bool `json:"hasFeed"`

		AbExperiments []UserAbExperiment `json:"abExperiments"`

		ImportMusicEnabled bool `json:"importMusicEnabled"`

		Status UserStatus `json:"status"`

		AccountAgeType string `json:"accountAgeType"`
	}

	UserSettings struct {
		SimpleUser

		Settings struct {
			PushOnRecommendations bool `json:"pushOnRecommendations"`

			MailOnRecommendations bool `json:"mailOnRecommendations"`

			PushOnNews bool `json:"pushOnNews"`

			PrivateAccount bool `json:"privateAccount"`

			MailOnNews bool `json:"mailOnNews"`

			MailOnPlaylistsUpdate bool `json:"mailOnPlaylistsUpdate"`

			PushOnPlaylistsUpdate bool `json:"pushOnPlaylistsUpdate"`

			PushOnNewMusic bool `json:"pushOnNewMusic"`

			FilterExplicitRecommendations bool `json:"filterExplicitRecommendations"`

			MailOnNewMusic bool `json:"mailOnNewMusic"`
		} `json:"settings"`
	}

	UserStatus struct {
		Editor bool `json:"editor"`

		Resident bool `json:"resident"`

		Moderator bool `json:"moderator"`

		Promouser bool `json:"promouser"`

		ServiceAccount bool `json:"serviceAccount"`
	}

	UserUpdateTime struct {
		MyMusicPane UnixTime `json:"myMusicPane"`

		Artists UnixTime `json:"artists"`

		SnippetFeed UnixTime `json:"snippetFeed"`

		AudioUpdatesFeed UnixTime `json:"audioUpdatesFeed"`

		Albums UnixTime `json:"albums"`

		CelebrityPlaylists UnixTime `json:"celebrityPlaylists"`

		Playlists UnixTime `json:"playlists"`
	}

	UserBilling struct {
		ComboAvailable bool `json:"comboAvailable"`

		SubscriptionRegion string `json:"subscriptionRegion"`

		ExpiryDate int `json:"expiryDate"`

		TrialAvailable bool `json:"trialAvailable"`

		Rank int `json:"rank"`

		State string `json:"state"`
	}

	UserCounts struct {
		Album int `json:"album"`

		Track int `json:"track"`

		Playlist int `json:"playlist"`

		Friend int `json:"friend"`

		Artist int `json:"artist"`
	}

	UserAbExperiment struct {
		ID string `json:"id"`

		Group int `json:"group"`
	}

	UserPlaylistSyncProgress struct {
		Ready int `json:"ready"`

		Total int `json:"total"`
	}
)
