package pexels_sdk_go

import (
	"context"
	"log"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

const (
	photoBaseURL      = "https://api.pexels.com/v1"
	videoBaseURL      = "https://api.pexels.com/videos"
	collectionBaseURL = "https://api.pexels.com/v1/collections"
)

var (
	validate = validator.New()
)

type option struct {
	debug bool
}

type OptFunc func(*option)

func WithDebug(debug bool) OptFunc {
	return func(o *option) {
		o.debug = debug
	}
}

func NewClient(apiKey string, opf ...OptFunc) *Client {
	option := &option{}
	for _, f := range opf {
		f(option)
	}

	c := &Client{
		option: option,
		apiKey: apiKey,
		rc:     resty.New(),
	}

	if c.option.debug {
		c.rc.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			return nil
		})

		c.rc.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
			log.Printf("url: %s", response.Request.URL)
			log.Printf("body: %s\n", response.Body())
			return nil
		})
	}

	return c
}

type Client struct {
	option *option
	apiKey string

	rc *resty.Client
}

func (c *Client) r(ctx context.Context) *resty.Request {
	return c.rc.R().
		SetContext(ctx).
		SetHeader("Authorization", c.apiKey)
}

func encode(v interface{}) (url.Values, error) {
	err := validate.Struct(v)
	if err != nil {
		return nil, err
	}
	return query.Values(v)
}

type Pagination struct {
	// The page number you are requesting. Default: 1
	Page int `url:"page,omitempty"`
	// The number of results you are requesting per page.
	// Default: 15 Max: 80
	PerPage int `url:"per_page,omitempty"`
}

type Cursor struct {
	// URL for the previous page of results, if applicable.
	PrevPage string `json:"prev_page"`
	// URL for the next page of results, if applicable.
	NextPage string `json:"next_page"`
}
