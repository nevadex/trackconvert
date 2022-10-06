package trackconvert

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/imroc/req/v3"
)

func searchTrackName(s string) ([]SearchResult, error) {
	resp, err := req.R().
		SetHeaders(headers).
		SetBody(strings.NewReader(`{"query":"` + strings.ReplaceAll(s, " ", "+") + `","context":{"client":{"hl":"en","gl":"US","clientName":"MWEB","clientVersion":"2.20220929.09.00"}}}`)).
		Post(ytSearchEndpoint)

	if err != nil || resp.StatusCode != 200 {
		return []SearchResult{}, errors.New("request failed: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []SearchResult{}, errors.New("reading failed: " + err.Error())
	}

	var queryContents []byte
	if _, t, _, _ := jsonparser.Get(body, "contents", "sectionListRenderer", "contents", "[2]"); t != jsonparser.NotExist {
		queryContents, _, _, err = jsonparser.Get(body, "contents", "sectionListRenderer", "contents", "[1]", "itemSectionRenderer", "contents")
	} else {
		queryContents, _, _, err = jsonparser.Get(body, "contents", "sectionListRenderer", "contents", "[0]", "itemSectionRenderer", "contents")
	}
	if err != nil {
		return []SearchResult{}, errors.New("raw parsing failed: " + err.Error())
	}

	var rs []SearchResult

	_, err = jsonparser.ArrayEach(queryContents, func(i []byte, dataType jsonparser.ValueType, _ int, err error) {
		if err != nil {
			return
		}
		r := SearchResult{}
		b, btype, _, err := jsonparser.Get(i, "compactVideoRenderer")
		if btype == jsonparser.NotExist {
			return
		}
		if err != nil {
			return
		}

		rawTimeB, vt, _, err := jsonparser.Get(b, "lengthText", "runs", "[0]")
		if vt == jsonparser.NotExist {
			return // ignore livestreams
		}
		if err != nil {
			return
		}

		r.Title, err = jsonparser.GetString(b, "title", "runs", "[0]", "text")
		if err != nil {
			return
		}

		r.VideoId, err = jsonparser.GetString(b, "videoId")
		if err != nil {
			return
		}

		r.Channel, err = jsonparser.GetString(b, "longBylineText", "runs", "[0]", "text")
		if err != nil {
			return
		}

		rawTime, err := jsonparser.GetString(rawTimeB, "text")
		if err != nil {
			return
		}
		r.Length, err = time.ParseDuration(strings.ReplaceAll(rawTime, ":", "m") + "s")
		if err != nil {
			return
		}

		rawViews, err := jsonparser.GetString(b, "viewCountText", "runs", "[0]", "text")
		if err != nil {
			return
		}
		r.Views, err = strconv.ParseInt(strings.ReplaceAll(strings.Split(rawViews, " ")[0], ",", ""), 10, 64)
		if err != nil {
			return
		}

		badge, err := jsonparser.GetString(b, "ownerBadges", "[0]", "metadataBadgeRenderer", "style")
		if badge == "BADGE_STYLE_TYPE_VERIFIED_ARTIST" && err == nil {
			r.VerifiedArtist = true
		} else {
			r.VerifiedArtist = false
		}

		rs = append(rs, r)
	})
	if err != nil {
		return []SearchResult{}, errors.New("array each failed: " + err.Error())
	}

	return rs, nil
}
