package artisanal

import (
	"context"
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-artisanal-integers-proxy/service"
	brooklyn_api "github.com/aaronland/go-brooklynintegers-api"
	"github.com/aaronland/go-uid"
	boltdb_pool "github.com/whosonfirst/go-whosonfirst-pool-boltdb"
	"strconv"
	"net/url"
)

func init() {
	ctx := context.Background()
	pr := NewArtisanalProvider()
	uid.RegisterProvider(ctx, "artisanal", pr)
}

type ArtisanalProvider struct {
	uid.Provider
	proxy artisanalinteger.Service
}

type ArtisanalUID struct {
	uid.UID
	// integer artisanalinteger.Integer
	integer int64
}

func NewArtisanalProvider() uid.Provider {
	pr := &ArtisanalProvider{}
	return pr
}

func (pr *ArtisanalProvider) Open(ctx context.Context, uri string) error {

	u, err := url.Parse(uri)

	if err != nil {
		return err
	}

	q := u.Query()

	str_min := q.Get("minimum")
	
	// artisanal:///?client=brooklyn&client=mission
	
	clients := make([]artisanalinteger.Client, 0)

	cl := brooklyn_api.NewAPIClient()
	clients = append(clients, cl)

	// pool, err := pool.NewPool(pool_uri)
	
	// please make me flags (see above)
	db := "integers.db"
	bucket := "integers"
	
	pool, err := boltdb_pool.NewBoltDBLIFOIntPool(db, bucket)
	
	if err != nil {
		return err
	}
	
	opts, err := service.DefaultProxyServiceOptions()

	if err != nil {
		return err
	}

	opts.Pool = pool

	min, err := strconv.Atoi(str_min)

	if err != nil {
		return err
	}
	
	opts.Minimum = min

	svc, err := service.NewProxyService(opts, clients...)

	if err != nil {
		return err
	}

	pr.proxy = svc
	return nil
}

func (pr *ArtisanalProvider) UID(...interface{}) (uid.UID, error) {

	i, err := pr.proxy.NextInt()

	if err != nil {
		return nil, err
	}

	return NewArtisanalUID(i)
}

func NewArtisanalUID(int int64) (uid.UID, error) {

	u := ArtisanalUID{
		integer: int,
	}

	return &u, nil
}

func (u *ArtisanalUID) String() string {

	return strconv.FormatInt(u.integer, 10)
}
