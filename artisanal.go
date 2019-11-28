package artisanal

import (
	"context"
	"errors"
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-artisanal-integers-proxy/service"
	"github.com/aaronland/go-uid"
	"github.com/aaronland/go-pool"
	"log"
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
	
	clients := make([]artisanalinteger.Client, 0)

	for _, cl_uri := range q["client"] {

		log.Println("WHAT", cl_uri)
		cl, err := artisanalinteger.NewClient(ctx, cl_uri)

		if err != nil {
			return err
		}
		
		clients = append(clients, cl)
	}
	
	if len(clients) == 0 {
		return errors.New("No artisanal integer clients defined")
	}
	
 	pool_uri := q.Get("pool")
	
	pl, err := pool.NewPool(ctx, pool_uri)
	
	if err != nil {
		return err
	}

	str_min := q.Get("minimum")	
	min, err := strconv.Atoi(str_min)

	if err != nil {
		return err
	}
	
	svc_opts, err := service.DefaultProxyServiceOptions()

	if err != nil {
		return err
	}

	svc_opts.Pool = pl	
	svc_opts.Minimum = min

	svc, err := service.NewProxyService(svc_opts, clients...)

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
