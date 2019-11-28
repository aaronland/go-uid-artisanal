package artisanal

import (
	"context"
	"testing"
	"github.com/aaronland/go-uid"
	_ "github.com/aaronland/go-brooklynintegers-api"	
)

func TestArtisanalProvider(t *testing.T) {

	ctx := context.Background()
	
	uri := "artisanal:///?pool=memory://&minimum=5&client=brooklynintegers://"
	
	pr, err := uid.NewProvider(ctx, uri)

	if err != nil {
		t.Fatal(err)
	}

	id, err := pr.UID()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(id)
}
