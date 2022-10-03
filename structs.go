package trackconvert

import (
	"strings"
	"time"
	"unicode"
)

func cleanString(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)
}

type searchResult struct {
	title          string
	videoId        string
	channel        string
	length         time.Duration
	verifiedArtist bool
}

type Track struct {
	id       string
	accuracy int
}

var headers = map[string]string{
	"Accept":       "*/*",
	"Content-Type": "application/json",
	"Host":         "www.youtube.com",
	"Referer":      "https://www.youtube.com/",
	"Origin":       "https://www.youtube.com",
	"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
}

var ytSearchEndpoint = "https://www.youtube.com/youtubei/v1/search?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
