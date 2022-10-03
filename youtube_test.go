package trackconvert

import (
	"strconv"
	"testing"
)

func TestSearchTrackName(t *testing.T) {
	rs, err := searchTrackName("J. Cole - Rise and Shine")

	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf(strconv.Itoa(len(rs)) + " results")
	for i := 0; i < len(rs); i++ {
		t.Logf(rs[i].videoId)
		t.Logf(rs[i].title)
		t.Logf(rs[i].channel)
		t.Logf(rs[i].length.String())
		t.Logf(strconv.FormatBool(rs[i].verifiedArtist))
		t.Logf("\n")
	}
}
