package pexels_sdk_go

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_SearchVideos(t *testing.T) {
	rsp, err := client.SearchVideos(context.Background(), &SearchVideosReq{
		Query:       "nature",
		Orientation: "",
		Size:        "",
		Locale:      "",
		Pagination: Pagination{
			PerPage: 1,
		},
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_PopularVideos(t *testing.T) {
	rsp, err := client.PopularVideos(context.Background(), &PopularVideosReq{
		MinWidth:    0,
		MinHeight:   0,
		MinDuration: 0,
		MaxDuration: 0,
		Pagination: Pagination{
			PerPage: 1,
		},
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetVideo(t *testing.T) {
	rsp, err := client.GetVideo(context.Background(), &GetVideoReq{ID: 2499611})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
