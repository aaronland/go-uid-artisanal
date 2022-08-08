package artisanal

import (
	"context"
	_ "github.com/aaronland/go-brooklynintegers-api"
	"github.com/aaronland/go-uid"
	"testing"
)

func TestArtisanalProvider(t *testing.T) {

	ctx := context.Background()

	uri := "artisanal://brooklynintegers"

	pr, err := uid.NewProvider(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create provider for %s, %v", uri, err)
	}

	_, err = pr.UID(ctx)

	if err != nil {
		t.Fatalf("Failed to create UID, %v", err)
	}

}
