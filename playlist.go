package govkm

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oklookat/govkm/schema"
)

// Получить плейлист по ID.
func (c Client) Playlist(ctx context.Context, id schema.ID) (*schema.Response[schema.PlaylistResponse], error) {
	return getEntityById[schema.PlaylistResponse](ctx, &c, "playlist", id)
}

// Получить треки в плейлисте.
func (c Client) PlaylistTracks(ctx context.Context, id schema.ID, limit, offset int) (*schema.Response[schema.TracksResponse], error) {
	params := getLimitOffset(limit, offset)
	return getAny[schema.TracksResponse](ctx, &c, params, "playlist", id.String(), "tracks/")
}

// Лайкнуть плейлист.
func (c Client) LikePlaylist(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return likeUnlikeEntity(ctx, &c, "playlist", id, true)
}

// Снять лайк с плейлиста.
func (c Client) UnlikePlaylist(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return likeUnlikeEntity(ctx, &c, "playlist", id, false)
}

// Создать плейлист.
func (c Client) CreatePlaylist(ctx context.Context, name string) (*schema.Response[schema.PlaylistResponse], error) {
	body := url.Values{}
	body.Set("name", name)
	return writeableAny[schema.PlaylistResponse](ctx, &c, http.MethodPost, body, "playlist/")
}

// Добавить трек в плейлист. Если трек уже будет в плейлисте, вернет nil плейлист.
func (c Client) AddTrackToPlaylist(ctx context.Context, id schema.ID, trackId schema.ID) (*schema.Response[schema.PlaylistResponse], error) {
	body := url.Values{}
	body.Set("file_id", trackId.String())
	return writeableAny[schema.PlaylistResponse](ctx, &c,
		http.MethodPut, body, "playlist", id.String(), "tracks/")
}

// Редактировать плейлист.
//
// name - новое или старое имя плейлиста.
//
// id - id плейлиста.
//
// tracks - id треков которые есть в плейлисте.
//
// 1. Вернет ошибку если в tracks будет трек которого нет в плейлисте.
//
// 2. Треки в плейлисте будут в порядке очередности, как в tracks.
// То есть так можно менять позицию треков в плейлисте.
//
// 3. Также можно удалить треки из плейлиста, если
// убрать ID нужных треков.
//
// 4. Если в tracks ничего не будет,
// будут удалены все треки из плейлиста.
func (c Client) EditPlaylist(
	ctx context.Context,
	name string,
	id schema.ID,
	tracks []schema.ID,
) (*schema.Response[schema.PlaylistResponse], error) {
	body := valuesFileId(tracks)
	body.Set("name", name)
	truncate := "false"
	if len(tracks) == 0 {
		truncate = "true"
	}
	body.Set("truncate", truncate)
	return writeableAny[schema.PlaylistResponse](ctx, &c, http.MethodPut, body, "playlist", id.String())
}

// Переименовать плейлист.
func (c Client) RenamePlaylist(ctx context.Context, id schema.ID, newName string) (*schema.Response[schema.PlaylistResponse], error) {
	body := url.Values{}
	body.Set("name", newName)
	body.Set("truncate", "false")
	return writeableAny[schema.PlaylistResponse](ctx, &c,
		http.MethodPut, body, "playlist", id.String())
}

// Удалить плейлист.
func (c Client) DeletePlaylist(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return writeableAny[any](ctx, &c,
		http.MethodDelete, nil, "playlist", id.String())
}

// Добавить все треки из плейлиста в новый плейлист.
//
// from - id плейлиста откуда взять треки.
//
// toName - имя нового плейлиста.
//
// Возвращает новый плейлист.
func (c Client) PlaylistTracksToPlaylist(
	ctx context.Context,
	from schema.ID,
	toName string,
) (*schema.Response[schema.PlaylistResponse], error) {
	body := url.Values{}
	body.Set("name", toName)
	return writeableAny[schema.PlaylistResponse](ctx, &c,
		http.MethodPost, body, "playlist", "playlist", from.String(), "/")
}
