package govkm

import (
	"context"
	"testing"
)

func TestAlbum(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.Album(ctx, _albumIds[0])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Album.APIID) == 0 {
		t.Fatal("len(resp.Data.Album.APIID) == 0")
	}
}

func TestLikeUnlikeAlbum(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _albumIds[0]

	_, err := cl.LikeAlbum(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	_, err = cl.UnlikeAlbum(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlbumRelevantPlaylsits(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.AlbumRelevantPlaylists(ctx, _albumIds[0], 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Playlists) == 0 {
		t.Fatal("len(resp.Data.Playlists) == 0")
	}
}
