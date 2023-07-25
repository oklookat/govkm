package govkm

import (
	"context"
	"testing"
)

func TestTrack(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.Track(ctx, _trackIds[0])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Track.APIID) == 0 {
		t.Fatal("len(resp.Data.Track.APIID) == 0")
	}
}

func TestTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.Tracks(ctx, _trackIds[:])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Tracks) == 0 {
		t.Fatal("len(resp.Data.Tracks) == 0")
	}
	if len(resp.Data.Tracks[0].APIID) == 0 {
		t.Fatal("len(resp.Data.Tracks[0].APIID) == 0")
	}
}

func TestLikeUnlikeTrack(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	id := _trackIds[0]

	_, err := cl.LikeTrack(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	_, err = cl.UnlikeTrack(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
}
