package govkm

import (
	"context"

	"github.com/oklookat/govkm/schema"
)

// Получить информацию о текущем пользователе.
func (c Client) UserInfo(ctx context.Context) (*schema.Response[schema.UserResponse], error) {
	return getAny[schema.UserResponse](ctx, &c, nil, "user", "info")
}

// Получить настройки текущего пользователя.
func (c Client) UserSettings(ctx context.Context) (*schema.Response[schema.UserSettingsResponse], error) {
	return getAny[schema.UserSettingsResponse](ctx, &c, nil, "user", "settings")
}

// Получить лайкнутых артистов.
func (c Client) LikedArtists(ctx context.Context, limit, offset int) (*schema.Response[schema.ArtistsResponse], error) {
	return getLiked[schema.ArtistsResponse](ctx, &c, "artists", limit, offset)
}

// Получить лайкнутые альбомы.
func (c Client) LikedAlbums(ctx context.Context, limit, offset int) (*schema.Response[schema.AlbumsResponse], error) {
	return getLiked[schema.AlbumsResponse](ctx, &c, "albums", limit, offset)
}

// Получить плейлисты в библиотеке текущего пользователя.
func (c Client) UserPlaylists(ctx context.Context, limit, offset int) (*schema.Response[schema.PlaylistsResponse], error) {
	params := getLimitOffset(limit, offset)
	return getAny[schema.PlaylistsResponse](ctx, &c, params, "user", "playlists/")
}

// Обертка над UserPlaylists, чтобы найти плейлист с лайкнутыми треками.
func (c Client) LikedTracks(ctx context.Context) (*schema.Playlist, error) {
	resp, err := c.UserPlaylists(ctx, 100, 0)
	if err != nil {
		return nil, err
	}
	for i := range resp.Data.Playlists {
		if resp.Data.Playlists[i].IsDefault {
			return &resp.Data.Playlists[i], err
		}
	}
	return nil, err
}
