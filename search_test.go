package govkm

import (
	"context"
	"strings"
	"testing"

	"github.com/oklookat/govkm/schema"
)

func TestSearch(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.Search(ctx, "сергей лазарев", 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Artists) == 0 {
		t.Fatal("len(resp.Data.Artists) == 0")
	}
}

func TestSearchSuggestion(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchSuggestion(ctx, "cardi")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Suggestions) == 0 {
		t.Fatal("len(resp.Data.Suggestions) == 0")
	}
}

func TestSearchTrack(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchTrack(ctx, "zhu", 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Tracks) == 0 {
		t.Fatal("len(resp.Data.Tracks) == 0")
	}
}

func TestSearchArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchArtist(ctx, "сергей лазарев", 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Artists) == 0 {
		t.Fatal("len(resp.Data.Artists) == 0")
	}
}

func TestSearchAlbum(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchAlbum(ctx, "сергей лазарев", 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Albums) == 0 {
		t.Fatal("len(resp.Data.Albums) == 0")
	}
}

func TestSearchPlaylist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchPlaylist(ctx, "музыка в машину", 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Playlists) == 0 {
		t.Fatal("len(resp.Data.Playlists) == 0")
	}
}

// For debug.
func TestFarmEntitiesIds(t *testing.T) {
	var (
		artistIds   []schema.ID
		albumIds    []schema.ID
		trackIds    []schema.ID
		playlistIds []schema.ID
	)

	ctx := context.Background()
	cl := getClient(t)
	resp, err := cl.Search(ctx, "сергей лазарев", 6, 0)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range resp.Data.Artists {
		artistIds = append(artistIds, e.APIID)
	}
	for _, e := range resp.Data.Albums {
		albumIds = append(albumIds, e.APIID)
	}
	for _, e := range resp.Data.Tracks {
		trackIds = append(trackIds, e.APIID)
	}
	for _, e := range resp.Data.Playlists {
		playlistIds = append(playlistIds, e.APIID)
	}

	printJoiner := func(entityType string, slice []schema.ID) {
		var elems []string
		for _, id := range slice {
			elems = append(elems, id.String())
		}
		println(entityType, ":", strings.Join(elems, ","))
	}

	printJoiner("artist", artistIds)
	printJoiner("album", albumIds)
	printJoiner("track", trackIds)
	printJoiner("playlist", playlistIds)
}
