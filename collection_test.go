package pexels_sdk_go

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_FeaturedCollections(t *testing.T) {
	rsp, err := client.FeaturedCollections(context.Background(), &FeaturedCollectionsReq{})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_MyCollections(t *testing.T) {
	rsp, err := client.MyCollections(context.Background(), &MyCollectionsReq{})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_CollectionMedia(t *testing.T) {
	rsp, err := client.CollectionMedia(context.Background(), &CollectionMediaReq{
		ID:         "iih3efu",
		Type:       "",
		Pagination: Pagination{},
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
