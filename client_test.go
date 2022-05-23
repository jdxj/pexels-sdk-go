package pexels_sdk_go

import (
	"fmt"
	"os"
	"testing"
)

var (
	client *Client
)

func TestMain(t *testing.M) {
	client = NewClient(apiKey, WithDebug(true))
	os.Exit(t.Run())
}

func TestEncode(t *testing.T) {
	sr := &SearchPhotosReq{
		Query:       "test",
		Orientation: "",
		Size:        "fff",
		Color:       "",
		Locale:      "",
		Pagination: Pagination{
			Page:    1,
			PerPage: 20,
		},
	}
	v, err := encode(sr)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Println(v.Encode())
}
