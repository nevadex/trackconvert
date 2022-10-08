package trackconvert

// Search for a song and return the most accurate result.
func ConvertSongAccurate(sd AccurateSongData) (Song, error) {
	rs, err := searchTrackName(makeTrackStr(cleanString(sd.Artist), cleanString(sd.Title)))
	if err != nil { return Song{}, err }

	r, ac := findBestResult(sd, rs)

	return Song{VideoId: r.VideoId, Accuracy: ac, SearchResult: r}, nil
}

// Search for a track, and return the first result.
// For optimal results, use ConvertSongAccurate().
func GetFirstResult(ar string, tr string) (SearchResult, error) {
	x, e := searchTrackName(makeTrackStr(ar, tr))
	if e == nil {
		return x[0], nil
	} else {
		return SearchResult{}, e
	}
}

// Search for a track, and return an array of results.
// For optimal results, use ConvertSongAccurate().
func GetAllResults(ar string, tr string) ([]SearchResult, error) {
	return searchTrackName(makeTrackStr(ar, tr))
}
