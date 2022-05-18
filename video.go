package pexels_sdk_go

import (
	"context"
	"fmt"
	"strconv"
)

type User struct {
	// The id of the videographer.
	ID uint64 `json:"id"`
	// The name of the videographer.
	Name string `json:"name"`
	// The URL of the videographer's Pexels profile.
	URL string `json:"url"`
}

type Quality string

const (
	HD Quality = "hd"
	SD Quality = "sd"
)

type VideoFile struct {
	// The id of the video_file.
	ID uint64 `json:"id"`
	// The video quality of the video_file.
	Quality Quality `json:"quality"`
	// The video format of the video_file.
	FileType string `json:"file_type"`
	// The width of the video_file in pixels.
	Width int `json:"width"`
	// The height of the video_file in pixels.
	Height int `json:"height"`
	// A link to where the video_file is hosted.
	Link string `json:"link"`
}

type VideoPicture struct {
	// The id of the video_picture.
	ID uint64 `json:"id"`
	// A link to the preview image.
	Picture string `json:"picture"`
	NR      int64  `json:"nr"`
}

type Video struct {
	// The id of the video.
	ID uint64 `json:"id"`
	// The real width of the video in pixels.
	Width int `json:"width"`
	// The real height of the video in pixels.
	Height int `json:"height"`
	// The Pexels URL where the video is located.
	URL string `json:"url"`
	// URL to a screenshot of the video.
	Image string `json:"image"`
	// The duration of the video in seconds.
	Duration int64 `json:"duration"`
	// The videographer who shot the video.
	User User `json:"user"`
	// An array of different sized versions of the video.
	VideoFiles []VideoFile `json:"video_files"`
	// An array of preview pictures of the video.
	VideoPictures []VideoPicture `json:"video_pictures"`
}

type SearchVideosReq struct {
	// The search query. Ocean, Tigers, Pears, etc.
	Query string `url:"query" validate:"required"`
	// Desired video orientation. The current supported
	// orientations are: landscape, portrait or square.
	Orientation Orientation `url:"orientation,omitempty"`
	// Minimum video size. The current supported sizes
	// are: large(4K), medium(Full HD) or small(HD).
	Size Size `url:"size,omitempty"`
	// The locale of the search you are performing.
	// The current supported locales are: 'en-US' 'pt-BR'
	// 'es-ES' 'ca-ES' 'de-DE' 'it-IT' 'fr-FR' 'sv-SE'
	// 'id-ID' 'pl-PL' 'ja-JP' 'zh-TW' 'zh-CN' 'ko-KR'
	// 'th-TH' 'nl-NL' 'hu-HU' 'vi-VN' 'cs-CZ' 'da-DK'
	// 'fi-FI' 'uk-UA' 'el-GR' 'ro-RO' 'nb-NO' 'sk-SK'
	// 'tr-TR' 'ru-RU'.
	Locale Locale `url:"locale,omitempty"`

	Pagination
}

type VideoList struct {
	// An array of Video objects.
	Videos []Video `json:"videos"`
	// The Pexels URL for the current search query.
	URL string `json:"url"`
	// The total number of results for the request.
	TotalResults int `json:"total_results"`

	Pagination
	Cursor
}

func (c *Client) SearchVideos(ctx context.Context, req *SearchVideosReq) (*VideoList, error) {
	v, err := encode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.r(ctx).
		SetQueryParamsFromValues(v).
		SetResult(&VideoList{}).
		Get(videoBaseURL + "/search")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%s", rsp.Status())
	}
	return rsp.Result().(*VideoList), nil
}

type PopularVideosReq struct {
	// The minimum width in pixels of the returned videos.
	MinWidth int `url:"min_width,omitempty"`
	// The minimum height in pixels of the returned videos.
	MinHeight int `url:"min_height,omitempty"`
	// The minimum duration in seconds of the returned videos.
	MinDuration int64 `url:"min_duration,omitempty"`
	// The maximum duration in seconds of the returned videos.
	MaxDuration int64 `url:"max_duration,omitempty"`

	Pagination
}

func (c *Client) PopularVideos(ctx context.Context, req *PopularVideosReq) (*VideoList, error) {
	v, err := encode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.r(ctx).
		SetQueryParamsFromValues(v).
		SetResult(&VideoList{}).
		Get(videoBaseURL + "/popular")
	if rsp.IsError() {
		return nil, fmt.Errorf("%s", rsp.Status())
	}
	return rsp.Result().(*VideoList), nil
}

type GetVideoReq struct {
	ID uint64 `url:"id" validate:"required"`
}

func (c *Client) GetVideo(ctx context.Context, req *GetVideoReq) (*Video, error) {
	_, err := encode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.r(ctx).
		SetPathParam("id", strconv.FormatUint(req.ID, 10)).
		SetResult(&Video{}).
		Get(videoBaseURL + "/videos/{id}")
	if rsp.IsError() {
		return nil, fmt.Errorf("%s", rsp.Status())
	}
	return rsp.Result().(*Video), nil
}
