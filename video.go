package pexels_sdk_go

import "context"

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
	// Desired video orientation. The current supported orientations are: landscape, portrait or square.
}

func (c *Client) SearchVideos(ctx context.Context) {}
