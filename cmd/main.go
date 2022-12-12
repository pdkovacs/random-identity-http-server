package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

const default_address = "0.0.0.0:8080"

type randomIdentityServer struct {
	identity string
}

func (randIdentServer *randomIdentityServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Hello! I am %s", randIdentServer.identity)))
}

func main() {
	rand.Seed(time.Now().Unix())
	addr := os.Getenv("SERVER_ADRESS")
	if addr == "" {
		addr = default_address
	}

	identity := namesgenerator.GetRandomName(0)
	handler := randomIdentityServer{
		identity: identity,
	}

	s := &http.Server{
		Addr:    addr,
		Handler: &handler,
	}
	log.Fatal(s.ListenAndServe())
}
