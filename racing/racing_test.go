package racing

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speeds of servers, returning url of fastest", func(t *testing.T) {
		// Mock HTTP servers
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("return an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(11 * time.Second)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Second)

		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
