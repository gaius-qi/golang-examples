package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/distribution/distribution/reference"
	"github.com/distribution/distribution/registry/client"
	"github.com/docker/distribution"
)

func main() {
	named, err := reference.ParseNamed("docker.io/library/busybox:latest")
	if err != nil {
		panic(err)
	}

	fmt.Println("named:", named.String())
	repo, err := client.NewRepository(named, "https://registry-1.docker.io", http.DefaultTransport)
	if err != nil {
		panic(err)
	}

	manifests, err := repo.Manifests(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	mf, err := manifests.Get(context.Background(), "", distribution.WithTag("latest"))
	if err != nil {
		panic(err)
	}

	fmt.Println("manifests:", mf)

	for _, v := range distribution.ManifestMediaTypes() {
		fmt.Println("media type:", v)
	}
}
