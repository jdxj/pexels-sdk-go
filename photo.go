package pexels_sdk_go

import (
	"context"
	"fmt"
)

type Orientation string

const (
	Landscape Orientation = "landscape"
	Portrait  Orientation = "portrait"
	Square    Orientation = "square"
)

type Size string

const (
	// Large 24MP
	Large Size = "large"
	// Medium 12MP
	Medium Size = "medium"
	// Small 4MP
	Small Size = "small"
)

type Color = string

const (
	Red       Color = "red"
	Orange    Color = "orange"
	Yellow    Color = "yellow"
	Green     Color = "green"
	Turquoise Color = "turquoise"
	Blue      Color = "blue"
	Violet    Color = "violet"
	Pink      Color = "pink"
	Brown     Color = "brown"
	Black     Color = "black"
	Gray      Color = "gray"
	White     Color = "white"
)

type Locale string

const (
	EN_US Locale = "en-US"
	PT_BR Locale = "pt-BR"
	ES_ES Locale = "es-ES"
	CA_ES Locale = "ca-ES"
	DE_DE Locale = "de-DE"
	IT_IT Locale = "it-IT"
	FR_FR Locale = "fr-FR"
	SV_SE Locale = "sv-SE"
	ID_ID Locale = "id-ID"
	PL_PL Locale = "pl-PL"
	JA_JP Locale = "ja-JP"
	ZH_TW Locale = "zh-TW"
	ZH_CN Locale = "zh-CN"
	KO_KR Locale = "ko-KR"
	TH_TH Locale = "th-TH"
	NL_NL Locale = "nl-NL"
	HU_HU Locale = "hu-HU"
	VI_VN Locale = "vi-VN"
	CS_CZ Locale = "cs-CZ"
	DA_DK Locale = "da-DK"
	FI_FI Locale = "fi-FI"
	UK_UA Locale = "uk-UA"
	EL_GR Locale = "el-GR"
	RO_RO Locale = "ro-RO"
	NB_NO Locale = "nb-NO"
	SK_SK Locale = "sk-SK"
	TR_TR Locale = "tr-TR"
	RU_RU Locale = "ru-RU"
)

type Src struct {
	// The image without any size changes.
	// It will be the same as the width and height attributes.
	Original string `json:"original"`
	// The image resized to W 940px X H 650px DPR 1.
	Large string `json:"large"`
	// The image resized W 940px X H 650px DPR 2.
	Large2x string `json:"large2x"`
	// The image scaled proportionally so that it's new height is 350px.
	Medium string `json:"medium"`
	// The image scaled proportionally so that it's new height is 130px.
	Small string `json:"small"`
	// The image cropped to W 800px X H 1200px.
	Portrait string `json:"portrait"`
	// The image cropped to W 1200px X H 627px.
	Landscape string `json:"landscape"`
	// The image cropped to W 280px X H 200px.
	Tiny string `json:"tiny"`
}

type Photo struct {
	// The id of the photo.
	ID uint64 `json:"id"`
	// The real width of the photo in pixels.
	Width int `json:"width"`
	// The real height of the photo in pixels.
	Height int `json:"height"`
	// The Pexels URL where the photo is located.
	URL string `json:"url"`
	// The name of the photographer who took the photo.
	Photographer string `json:"photographer"`
	// The URL of the photographer's Pexels profile.
	PhotographerURL string `json:"photographer_url"`
	// The id of the photographer.
	PhotographerID uint64 `json:"photographer_id"`
	// The average color of the photo. Useful for a
	// placeholder while the image loads.
	AvgColor string `json:"avg_color"`
	// An assortment of different image sizes that can
	// be used to display this Photo.
	Src   Src  `json:"src"`
	Liked bool `json:"liked"`
	// Text description of the photo for use in the alt attribute.
	Alt string `json:"alt"`
}

type PhotoList struct {
	// An array of Photo objects.
	Photos []Photo `json:"photos"`
	// The total number of results for the request.
	TotalResults int `json:"total_results"`

	Pagination
	Cursor
}

type SearchReq struct {
	// The search query. Ocean, Tigers, Pears, etc.
	Query string `url:"query" validate:"required"`
	// Desired photo orientation. The current supported
	// orientations are: landscape, portrait or square.
	Orientation Orientation `url:"orientation,omitempty"`
	// Minimum photo size. The current supported sizes
	// are: large(24MP), medium(12MP) or small(4MP).
	Size Size `url:"size,omitempty"`
	// Desired photo color. Supported colors: red, orange,
	// yellow, green, turquoise, blue, violet, pink, brown,
	// black, gray, white or any hexidecimal color code
	// (eg. #ffffff).
	Color Color `url:"color,omitempty"`
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

func (c *Client) Search(ctx context.Context, req *SearchReq) (*PhotoList, error) {
	v, err := encode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.r(ctx).
		SetQueryParamsFromValues(v).
		SetResult(&PhotoList{}).
		Get(baseURL + "/search")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%s", rsp.Status())
	}
	return rsp.Result().(*PhotoList), nil
}

type CuratedReq struct {
	Pagination
}

func (c *Client) Curated(ctx context.Context, req *CuratedReq) (*PhotoList, error) {
	v, err := encode(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.r(ctx).
		SetQueryParamsFromValues(v).
		SetResult(&PhotoList{}).
		Get(baseURL + "/curated")
	if rsp.IsError() {
		return nil, fmt.Errorf("%s", rsp.Status())
	}
	return rsp.Result().(*PhotoList), nil
}
