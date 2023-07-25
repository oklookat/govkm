package govkm

import (
	"context"

	"github.com/oklookat/govkm/schema"
)

// Поиск всего.
func (c Client) Search(ctx context.Context, query string, limit, offset int) (*schema.Response[schema.SearchResult], error) {
	return searchEntity[schema.SearchResult](ctx, &c, query, "", limit, offset)
}

// Получить подсказки по поиску.
func (c Client) SearchSuggestion(ctx context.Context, query string) (*schema.Response[schema.SearchSuggestion], error) {
	return searchEntity[schema.SearchSuggestion](ctx, &c, query, "suggestion", 0, 0)
}

// Поиск трека.
func (c Client) SearchTrack(ctx context.Context, query string, limit, offset int) (*schema.Response[schema.TracksResponse], error) {
	return searchEntity[schema.TracksResponse](ctx, &c, query, "track", limit, offset)
}

// Поиск артиста.
func (c Client) SearchArtist(ctx context.Context, query string, limit, offset int) (*schema.Response[schema.ArtistsResponse], error) {
	return searchEntity[schema.ArtistsResponse](ctx, &c, query, "artist", limit, offset)
}

// Поиск альбома.
func (c Client) SearchAlbum(ctx context.Context, query string, limit, offset int) (*schema.Response[schema.AlbumsResponse], error) {
	return searchEntity[schema.AlbumsResponse](ctx, &c, query, "album", limit, offset)
}

// Поиск плейлиста.
func (c Client) SearchPlaylist(ctx context.Context, query string, limit, offset int) (*schema.Response[schema.PlaylistsResponse], error) {
	return searchEntity[schema.PlaylistsResponse](ctx, &c, query, "playlist", limit, offset)
}

// TODO SearchPodcast
