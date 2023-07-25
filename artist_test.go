package govkm

import (
	"context"
	"testing"

	"github.com/oklookat/govkm/schema"
)

func TestArtist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.Artist(ctx, _artistIds[0])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Artist.APIID) == 0 {
		t.Fatal("len(resp.Data.Artist.APIID) == 0")
	}
}

func TestLikeUnlikeArtist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _artistIds[0]

	_, err := cl.LikeArtist(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	_, err = cl.UnlikeArtist(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestArtistTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _artistIds[0]

	resp, err := cl.ArtistTracks(ctx, id, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Tracks) == 0 {
		t.Fatal("len(resp.Data.Tracks) == 0")
	}
}

func TestArtistSingleTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _artistIds[0]

	resp, err := cl.ArtistSingleTracks(ctx, id, 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Tracks) == 0 {
		t.Fatal("len(resp.Data.Tracks) == 0")
	}
}

func TestArtistAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _artistIds[0]

	resp, err := cl.ArtistAlbums(ctx, id, []schema.AlbumType{
		schema.AlbumTypeAlbum,
		schema.AlbumTypeEp,
	}, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Albums) == 0 {
		t.Fatal("len(resp.Data.Albums) == 0")
	}
}

func TestArtistFeaturingAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := schema.ID("13178539")

	resp, err := cl.ArtistFeaturingAlbums(ctx, id, 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Albums) == 0 {
		t.Fatal("len(resp.Data.Albums) == 0")
	}
}

func TestArtistPlaylists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _artistIds[0]

	resp, err := cl.ArtistPlaylists(ctx, id, 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Playlists) == 0 {
		t.Fatal("len(resp.Data.Playlists) == 0")
	}
}

func TestRelevantArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _artistIds[0]

	resp, err := cl.RelevantArtists(ctx, id, 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Artists) == 0 {
		t.Fatal("len(resp.Data.Artists) == 0")
	}
}
