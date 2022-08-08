package artisanal

import (
	"context"
	"fmt"
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-uid"
	"net/url"
)

const ARTISANAL_SCHEME string = "artisanal"

func init() {
	ctx := context.Background()
	uid.RegisterProvider(ctx, ARTISANAL_SCHEME, NewArtisanalProvider)
}

type ArtisanalProvider struct {
	uid.Provider
	client artisanalinteger.Client
}

type ArtisanalUID struct {
	uid.UID
	id int64
}

func NewArtisanalProvider(ctx context.Context, uri string) (uid.Provider, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse client, %w", err)
	}

	client_uri := fmt.Sprintf("%s://", u.Host)

	client, err := artisanalinteger.NewClient(ctx, client_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create artisanal integer client, %w", err)
	}

	pr := &ArtisanalProvider{
		client: client,
	}

	return pr, nil
}

func (pr *ArtisanalProvider) UID(ctx context.Context, args ...interface{}) (uid.UID, error) {
	return NewArtisanalUID(ctx, pr.client)
}

func NewArtisanalUID(ctx context.Context, args ...interface{}) (uid.UID, error) {

	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid arguments")
	}

	cl, ok := args[0].(artisanalinteger.Client)

	if !ok {
		return nil, fmt.Errorf("Invalid client")
	}

	i, err := cl.NextInt()

	if err != nil {
		return nil, fmt.Errorf("Failed to create new integerm %w", err)
	}

	u := &ArtisanalUID{
		id: i,
	}

	return u, nil
}

func (u *ArtisanalUID) Value() any {
	return u.id
}

func (u *ArtisanalUID) String() string {
	return fmt.Sprintf("%v", u.Value())
}
