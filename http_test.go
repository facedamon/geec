package geec

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

var dbs = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestHttp(t *testing.T) {
	NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SLowDB] search key", key)
			if v, ok := dbs[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "127.0.0.1:9999"
	peers := NewHTTPPool(addr)
	log.Println("geec is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
