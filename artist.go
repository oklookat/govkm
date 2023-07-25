package govkm

import (
	"context"

	"github.com/oklookat/govkm/schema"
)

// Получить артиста по ID.
func (c Client) Artist(ctx context.Context, id schema.ID) (*schema.Response[schema.ArtistResponse], error) {
	return getEntityById[schema.ArtistResponse](ctx, &c, "artist", id)
}

// Лайкнуть артиста.
func (c Client) LikeArtist(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return likeUnlikeEntity(ctx, &c, "artist", id, true)
}

// Снять лайк с артиста.
func (c Client) UnlikeArtist(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return likeUnlikeEntity(ctx, &c, "artist", id, false)
}

// Получить все треки артиста. Сначала идут популярные.
func (c Client) ArtistTracks(ctx context.Context, id schema.ID, limit, offset int) (*schema.Response[schema.TracksResponse], error) {
	params := getLimitOffset(limit, offset)
	return getAny[schema.TracksResponse](ctx, &c, params, "artist", id.String(), "tracks/")
}

// Получить синглы артиста.
func (c Client) ArtistSingleTracks(ctx context.Context, id schema.ID, limit, offset int) (*schema.Response[schema.TracksResponse], error) {
	params := getLimitOffset(limit, offset)
	return getAny[schema.TracksResponse](ctx, &c, params, "artist", id.String(), "single_tracks/")
}

// Получить альбомы.
//
// Если в albumTypes ничего нет, будут получены все типы альбомов.
func (c Client) ArtistAlbums(ctx context.Context, id schema.ID, albumTypes []schema.AlbumType, limit, offset int) (*schema.Response[schema.AlbumsResponse], error) {
	params := getLimitOffset(limit, offset)
	for _, at := range albumTypes {
		params.Add("type", at.String())
	}
	return getAny[schema.AlbumsResponse](ctx, &c, params, "artist", id.String(), "albums/")
}

// Получить альбомы в которых артист принял участие.
func (c Client) ArtistFeaturingAlbums(ctx context.Context, id schema.ID, limit, offset int) (*schema.Response[schema.AlbumsResponse], error) {
	params := getLimitOffset(limit, offset)
	return getAny[schema.AlbumsResponse](ctx, &c, params, "artist", id.String(), "album", "featuring/")
}

// Получить плейлисты (от редакции?) в которых присутствует артист.
func (c Client) ArtistPlaylists(ctx context.Context, id schema.ID, limit, offset int) (*schema.Response[schema.PlaylistsResponse], error) {
	params := getLimitOffset(limit, offset)
	return getAny[schema.PlaylistsResponse](ctx, &c, params, "artist", id.String(), "playlists/")
}

// Получить похожих артистов.
func (c Client) RelevantArtists(ctx context.Context, id schema.ID, limit, offset int) (*schema.Response[schema.SimpleArtistsResponse], error) {
	params := getLimitOffset(limit, offset)
	return getAny[schema.SimpleArtistsResponse](ctx, &c, params, "artist", id.String(), "relevant_artists/")
}
