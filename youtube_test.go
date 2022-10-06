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
		t.Logf(rs[i].VideoId)
		t.Logf(rs[i].Title)
		t.Logf(rs[i].Channel)
		t.Logf(rs[i].Length.String())
		t.Logf(strconv.FormatInt(rs[i].Views, 10))
		t.Logf(strconv.FormatBool(rs[i].VerifiedArtist))
		t.Logf("\n")
	}
}
