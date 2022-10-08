package trackconvert

import (
	"math"
	"sort"
	"strings"
	"time"
)

// simple algorithm to find the closest yt video for a specific song
func findBestResult(sd AccurateSongData, sr []SearchResult) (SearchResult, int) {
	for index, i := range sr {
		a := 0
		if strings.Contains(strings.ToLower(i.Title), "remix") || strings.Contains(strings.ToLower(i.Title), "mashup") || strings.Contains(strings.ToLower(i.Title), "cover") || strings.Contains(strings.ToLower(i.Title), "remaster") {
			a += 15
		}

		if math.Abs(float64(i.Length-sd.Length)) > float64(time.Second*3) {
			a += 13
		}

		if i.Channel != sd.Artist {
			if strings.Contains(i.Channel, sd.Artist) {
				a += 5
			} else {
				a += 11
			}
		}

		if i.Title != sd.Title {
			if strings.Contains(strings.ToLower(i.Title), strings.ToLower(sd.Title)) && !strings.Contains(strings.ToLower(i.Title), "remix") {
				a += 3
			} else {
				a += 9
			}
		}

		if !i.VerifiedArtist {
			a += 7
		}
		sr[index].accuracy = a

		// debug
		//println(i.VideoId + "_" + i.Channel + "_" + i.Title + "_" + strconv.Itoa(a))
	}

	var bestSR []SearchResult
	bestAcc := 64 // worst accuracy + 1

	for _, i := range sr {
		if i.accuracy < bestAcc {
			bestAcc = i.accuracy
			bestSR = []SearchResult{}
			bestSR = append(bestSR, i)
		} else if i.accuracy == bestAcc {
			bestSR = append(bestSR, i)
		} else {
			continue
		}
	}

	sort.SliceStable(bestSR, func(i, j int) bool {
		return bestSR[i].Views > bestSR[j].Views
	})

	return bestSR[0], bestSR[0].accuracy
}
