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

// Search a track name, and return an array of results.
// If no results are returned from YT, there will be no error.
func searchTrackName(s string) ([]searchResult, error) {
	resp, err := req.R().
		SetHeaders(headers).
		SetBody(strings.NewReader(`{"query":"` + strings.ReplaceAll(s, " ", "+") + `","context":{"client":{"hl":"en","gl":"US","clientName":"MWEB","clientVersion":"2.20220929.09.00"}}}`)).
		Post(ytSearchEndpoint)

	if err != nil || resp.StatusCode != 200 { return []searchResult{}, errors.New("request failed: " + err.Error()) }

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { return []searchResult{}, errors.New("reading failed: " + err.Error()) }

	var queryContents []byte
	if _, t, _, _ := jsonparser.Get(body, "contents", "sectionListRenderer", "contents", "[2]"); t != jsonparser.NotExist {
		queryContents, _, _, err = jsonparser.Get(body, "contents", "sectionListRenderer", "contents", "[1]", "itemSectionRenderer", "contents")
	} else {
		queryContents, _, _, err = jsonparser.Get(body, "contents", "sectionListRenderer", "contents", "[0]", "itemSectionRenderer", "contents")
	}
	if err != nil { return []searchResult{}, errors.New("raw parsing failed: " + err.Error()) }
	
	var rs []searchResult
	for i := 0; i < 30; i++ {
		r := searchResult{}
		b, btype, _, err := jsonparser.Get(queryContents, "[" + strconv.Itoa(i) + "]", "compactVideoRenderer")
		if btype == jsonparser.NotExist {
			/*if _, sft, _, err := jsonparser.Get(queryContents, "[" + strconv.Itoa(i+1) + "]", "compactVideoRenderer"); i == 0 && sft != jsonparser.NotExist && err == nil {
				return []searchResult{}, errors.New("parsing failed: no results")
			} else */if i > 20 {
				break
			} else {
				continue
			}
		}
		if err != nil { return []searchResult{}, errors.New("parsing failed: " + err.Error()) }

		r.title, err = jsonparser.GetString(b, "title", "runs", "[0]", "text")
		if err != nil { return []searchResult{}, errors.New("parsing failed [title]: " + err.Error()) }

		r.videoId, err = jsonparser.GetString(b, "videoId")
		if err != nil { return []searchResult{}, errors.New("parsing failed [videoId]: " + err.Error()) }

		r.channel, err = jsonparser.GetString(b, "longBylineText", "runs", "[0]", "text")
		if err != nil { return []searchResult{}, errors.New("parsing failed [channel]: " + err.Error()) }

		rawTime, err := jsonparser.GetString(b, "lengthText", "runs", "[0]", "text")
		if err != nil { return []searchResult{}, errors.New("parsing failed [length.rawTime]: " + err.Error()) }
		r.length, err = time.ParseDuration(strings.ReplaceAll(rawTime, ":", "m") + "s")
		if err != nil { return []searchResult{}, errors.New("parsing failed [length]: " + err.Error()) }

		badge, err := jsonparser.GetString(b, "ownerBadges", "[0]", "metadataBadgeRenderer", "style")
		if badge == "BADGE_STYLE_TYPE_VERIFIED_ARTIST" && err == nil {
			r.verifiedArtist = true
		} else {
			r.verifiedArtist = false
		}

		rs = append(rs, r)
	}

	return rs, nil
}
