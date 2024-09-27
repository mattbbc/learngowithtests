package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	// Something that implements Fetch() is a Store
	Fetch(ctx context.Context) (string, error)
}

// http.HandlerFunc is a type that has a ServeHTTP method attached
// so whatever this returns will have a ServeHTTP() thingy
//
// Also because the type implements a ServeHTTP method this means that
// http.HandlerFunc adheres to the Handler interface https://pkg.go.dev/net/http#Handler
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
