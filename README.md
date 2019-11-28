# go-uid-artisanal

Work in progress.

## Example

```
package main

import (
	"context"
	_ "github.com/aaronland/go-brooklynintegers-api"
	"github.com/aaronland/go-uid"
	"log"
)

func main() {

	ctx := context.Background()

	opts := &ArtisanalProviderURIOptions{
		Pool:    "memory://",
		Minimum: 5,
		Clients: []string{
			"brooklynintegers://",
		},
	}

	uri, _ := NewArtisanalProviderURI(opts)
	str_uri := uri.String()

	pr, _ := uid.NewProvider(ctx, str_uri)
	id, _ := pr.UID()

	log.Println(id.String())
}

```

## See also

* https://github.com/aaronland/go-uid