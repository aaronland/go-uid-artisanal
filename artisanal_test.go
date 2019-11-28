package artisanal

import (
	"context"
	_ "github.com/aaronland/go-brooklynintegers-api"
	_ "github.com/aaronland/go-missionintegers-api"
	"github.com/aaronland/go-uid"
	"testing"
)

func TestArtisanalProvider(t *testing.T) {

	ctx := context.Background()

	opts := &ArtisanalProviderURIOptions{
		Pool:    "memory://",
		Minimum: 5,
		Clients: []string{
			"brooklynintegers://",
			"missionintegers://",
		},
	}

	uri, err := NewArtisanalProviderURI(opts)

	if err != nil {
		t.Fatal(err)
	}

	str_uri := uri.String()

	t.Log(str_uri)

	pr, err := uid.NewProvider(ctx, str_uri)

	if err != nil {
		t.Fatal(err)
	}

	id, err := pr.UID()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(id)
}
