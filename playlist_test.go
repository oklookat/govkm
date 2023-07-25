package govkm

import (
	"context"
	"testing"

	"github.com/oklookat/govkm/schema"
)

func TestPlaylist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.Playlist(ctx, _playlistIds[0])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Playlist.APIID) == 0 {
		t.Fatal("len(resp.Data.Playlist.APIID) == 0")
	}
}

func TestPlaylistTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.PlaylistTracks(ctx, _playlistIds[0], 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Tracks) == 0 {
		t.Fatal("len(resp.Data.Tracks) == 0")
	}
}

func TestLikeUnlikePlaylist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _playlistIds[0]

	_, err := cl.LikePlaylist(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	_, err = cl.UnlikePlaylist(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPlaylistCRUD(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// Create.
	plResp, err := cl.CreatePlaylist(ctx, "THE CRUD")
	if err != nil {
		t.Fatal(err)
	}
	if plResp.Data.Playlist == nil {
		t.Fatal("nil playlist")
	}
	pl := plResp.Data.Playlist

	// Add tracks.
	addTracks := func(id schema.ID) {
		plResp, err = cl.AddTrackToPlaylist(ctx, pl.APIID, id)
		if err != nil {
			t.Fatal(err)
		}
		if plResp.Data.Playlist == nil {
			t.Fatal("nil playlist")
		}
		pl = plResp.Data.Playlist
	}
	addTracks(_trackIds[0])

	// Playlist tracks to new playlist.
	newPl, err := cl.PlaylistTracksToPlaylist(ctx, pl.APIID, "THE THE THE")
	if err != nil {
		t.Fatal(err)
	}
	if newPl.Data.Playlist == nil {
		t.Fatal("nil playlist")
	}
	_, err = cl.DeletePlaylist(ctx, newPl.Data.Playlist.APIID)
	if err != nil {
		t.Fatal(err)
	}

	// Edit, delete all tracks.
	plResp, err = cl.EditPlaylist(ctx, "THE CROOK", pl.APIID, []schema.ID{})
	if err != nil {
		t.Fatal(err)
	}
	if plResp.Data.Playlist == nil {
		t.Fatal("nil playlist")
	}
	pl = plResp.Data.Playlist

	// Rename.
	plResp, err = cl.RenamePlaylist(ctx, pl.APIID, "THE CURE")
	if err != nil {
		t.Fatal(err)
	}
	if plResp.Data.Playlist == nil {
		t.Fatal("nil playlist")
	}
	pl = plResp.Data.Playlist

	// Delete.
	_, err = cl.DeletePlaylist(ctx, pl.APIID)
	if err != nil {
		t.Fatal(err)
	}
}
