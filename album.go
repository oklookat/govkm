package govkm

import (
	"context"

	"github.com/oklookat/govkm/schema"
)

// Получить альбом по ID.
func (c Client) Album(ctx context.Context, id schema.ID) (*schema.Response[schema.AlbumResponse], error) {
	return getEntityById[schema.AlbumResponse](ctx, &c, "album", id)
}

// Лайкнуть альбом.
func (c Client) LikeAlbum(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return likeUnlikeEntity(ctx, &c, "album", id, true)
}

// Снять лайк с альбома.
func (c Client) UnlikeAlbum(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return likeUnlikeEntity(ctx, &c, "album", id, false)
}

// Плейлисты (от редакции?) связанные с альбомом, артистом.
func (c Client) AlbumRelevantPlaylists(ctx context.Context, id schema.ID, limit, offset int) (*schema.Response[schema.PlaylistsResponse], error) {
	params := getLimitOffset(limit, offset)
	return getAny[schema.PlaylistsResponse](ctx, &c, params, "album", id.String(), "relevant", "playlists/")
}
