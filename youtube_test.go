package trackconvert

import (
	"strconv"
	"testing"
	"time"
)

func TestSearchTrackName(t *testing.T) {
	rs, err := searchTrackName("J. Cole - Rise and Shine")

	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf(strconv.Itoa(len(rs)) + " results")
	for i := 0; i < len(rs); i++ {
		t.Logf(rs[i].VideoId)
		t.Logf(rs[i].Title)
		t.Logf(rs[i].Channel)
		t.Logf(rs[i].Length.String())
		t.Logf(strconv.FormatInt(rs[i].Views, 10))
		t.Logf(strconv.FormatBool(rs[i].VerifiedArtist))
		t.Logf("\n")
	}
}

func TestAccurateTrack(t *testing.T) {
	sd := AccurateSongData{
		Title:  "Rise and Shine",
		Artist: "J. Cole",
	}
	sd.Length, _ = time.ParseDuration("4m35s")

	rs, err := searchTrackName(makeTrackStr(cleanString(sd.Artist), cleanString(sd.Title)))
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf(strconv.Itoa(len(rs)) + " results\n")
	r, ac := findBestResult(sd, rs)

	t.Logf(r.VideoId)
	t.Logf(r.Title)
	t.Logf(r.Channel)
	t.Logf(r.Length.String())
	t.Logf(strconv.FormatInt(r.Views, 10))
	t.Logf(strconv.FormatBool(r.VerifiedArtist))
	t.Logf("obj ac: " + strconv.Itoa(r.accuracy) + " | out ac: " + strconv.Itoa(ac))
	t.Logf("\n")
}
