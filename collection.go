package pexels_sdk_go

import (
	"context"
	"fmt"
)

type Collection struct {
	// The id of the collection.
	ID string `json:"id"`
	// The name of the collection.
	Title string `json:"title"`
	// The description of the collection.
	Description string `json:"description"`
	// Whether or not the collection is marked as private.
	Private bool `json:"private"`
	// The total number of media included in this collection.
	MediaCount int `json:"media_count"`
	// The total number of photos included in this collection.
	PhotosCount int `json:"photos_count"`
	// The total number of videos included in this collection.
	VideosCount int `json:"videos_count"`
}

type FeaturedCollectionsReq struct {
	Pagination
}

type CollectionList struct {
	// An array of Collection objects.
	Collections []Collection `json:"collections"`
	// The total number of results for the request.
	TotalResults int `json:"total_results"`

	Pagination
	Cursor
}

// FeaturedCollections This endpoint returns all featured
// collections on Pexels.
func (c *Client) FeaturedCollections(ctx context.Context, req *FeaturedCollectionsReq) (*CollectionList, error) {
	v, err := encode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.r(ctx).
		SetQueryParamsFromValues(v).
		SetResult(&CollectionList{}).
		Get(collectionBaseURL + "/featured")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%s", rsp.Status())
	}
	return rsp.Result().(*CollectionList), nil
}

type MyCollectionsReq struct {
	Pagination
}

// MyCollections This endpoint returns all of your collections.
func (c *Client) MyCollections(ctx context.Context, req *MyCollectionsReq) (*CollectionList, error) {
	v, err := encode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.r(ctx).
		SetQueryParamsFromValues(v).
		SetResult(&CollectionList{}).
		Get(collectionBaseURL)
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%s", rsp.Status())
	}
	return rsp.Result().(*CollectionList), nil
}

type Type string

const (
	Photos Type = "photos"
	Videos Type = "videos"
)

type CollectionMediaReq struct {
	// Collection id
	ID string `url:"-" validate:"required"`
	// The type of media you are requesting.
	// If not given or if given with an invalid value,
	// all media will be returned. Supported values
	// are photos and videos.
	Type Type `url:"type,omitempty"`

	Pagination
}

type Media struct {
	Type Type `json:"type"`
	// The id of the photo/video.
	ID uint64 `json:"id"`
	// The real width of the photo/video in pixels.
	Width int `json:"width"`
	// The real height of the photo/video in pixels.
	Height int `json:"height"`
	// The Pexels URL where the photo/video is located.
	URL string `json:"url"`

	Photo
	Video
}

type MediaList struct {
	// The id of the collection you are requesting.
	ID string `json:"id"`
	// An array of media objects. Each object has
	// an extra type attribute to indicate the
	// type of object.
	Medias []Media `json:"media"`
	// The total number of results for the request.
	TotalResults int `json:"total_results"`

	Pagination
	Cursor
}

// CollectionMedia This endpoint returns all the
// media (photos and videos) within a single collection.
// You can filter to only receive photos or videos using
// the type parameter.
func (c *Client) CollectionMedia(ctx context.Context, req *CollectionMediaReq) (*MediaList, error) {
	v, err := encode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.r(ctx).
		SetPathParam("id", req.ID).
		SetQueryParamsFromValues(v).
		SetResult(&MediaList{}).
		Get(collectionBaseURL + "/{id}")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%s", rsp.Status())
	}
	return rsp.Result().(*MediaList), nil
}
