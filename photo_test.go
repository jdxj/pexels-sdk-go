package pexels_sdk_go

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_Search(t *testing.T) {
	pl, err := client.SearchPhotos(context.Background(), &SearchPhotosReq{
		Query:       "apple",
		Orientation: "",
		Size:        "",
		Color:       "",
		Locale:      "",
		Pagination: Pagination{
			PerPage: 1,
		},
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", pl)
}

func TestClient_Curated(t *testing.T) {
	pl, err := client.CuratedPhotos(context.Background(), &CuratedReq{
		Pagination: Pagination{
			Page:    0,
			PerPage: 1,
		},
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", pl)
}

func TestClient_GetPhoto(t *testing.T) {
	rsp, err := client.GetPhoto(context.Background(), &GetPhotoReq{ID: 2014422})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
