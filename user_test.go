package govkm

import (
	"context"
	"testing"
)

func TestUserInfo(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.UserInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.User.APIID) == 0 {
		t.Fatal("len(resp.Data.User.APIID) == 0")
	}
}

func TestUserSettings(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.UserSettings(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.User.APIID) == 0 {
		t.Fatal("len(resp.Data.User.APIID) == 0")
	}
}

func TestLikedArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.LikedArtists(ctx, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLikedAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	als, err := cl.LikedAlbums(ctx, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	println(len(als.Data.Albums))
}

func TestLikedTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	pl, err := cl.LikedTracks(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if pl == nil {
		t.Fatal("liked tracks playlist cannot be nil")
	}
}

func TestUserPlaylists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	pls, err := cl.UserPlaylists(ctx, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	println(len(pls.Data.Playlists))
}
