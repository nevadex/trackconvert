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

func makeTrackStr(ar string, tr string) string {
	return cleanString(ar) + " - " + cleanString(tr)
}

type SearchResult struct {
	Title          string
	VideoId        string
	Channel        string
	Length         time.Duration
	Views          int64
	VerifiedArtist bool
	accuracy       int
}

type AccurateSongData struct {
	Title  string
	Artist string
	Length time.Duration
}

type Song struct {
	VideoId      string
	Accuracy     int
	SearchResult SearchResult
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
