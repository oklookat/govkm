package govkm

import (
	"context"

	"github.com/oklookat/govkm/schema"
)

// Получить трек по ID.
func (c Client) Track(ctx context.Context, id schema.ID) (*schema.Response[schema.TrackResponse], error) {
	return getEntityById[schema.TrackResponse](ctx, &c, "track", id)
}

// Получить треки по ID.
func (c Client) Tracks(ctx context.Context, id []schema.ID) (*schema.Response[schema.TracksResponse], error) {
	params := valuesFileId(id)
	return getAny[schema.TracksResponse](ctx, &c, params, "tracks/")
}

// Лайкнуть трек.
func (c Client) LikeTrack(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return likeUnlikeEntity(ctx, &c, "track", id, true)
}

// Снять лайки с трека.
func (c Client) UnlikeTrack(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	return likeUnlikeEntity(ctx, &c, "track", id, false)
}
