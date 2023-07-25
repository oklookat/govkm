package govkm

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/oklookat/govkm/schema"
)

var (
	_artistIds = [4]schema.ID{
		"13449406",
		"19858139",
		"16985090",
		"16941569",
	}
	_albumIds = [4]schema.ID{
		"23212224",
		"23145074",
		"22731308",
		"22731801",
	}
	_trackIds = [4]schema.ID{
		"2d6e5e50ee1c2480bea4fe093be19fde",
		"1b5d1f4fe596e69563fd5d9629dbe286",
		"5ad6547c8acbfde7b28f33fce2e4abcd",
		"c0644ba88b54fa86341c503446ec1acc",
	}
	_playlistIds = [4]schema.ID{
		"32209597",
		"817551217",
		"241811439",
		"32307949",
	}
)

func getClient(t *testing.T) *Client {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}

	cl, err := New(os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}
	cl.Http.SetRateLimit(1, time.Second)

	return cl
}
