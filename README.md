# TrackConvert
A light golang package for querying youtube and finding a song based on results from another application.
Example use case: Getting songs from spotify, and making a youtube playlist with the spotify songs.
## Features

 - Functional youtube query without API key, returning minimal results to minimize web traffic
 - Simple golang API for easy integration
 - Algorithm to find the most accurate version of a song on youtube
## Quick start
```bash
go get github.com/nevadex/trackconvert
```
### Usage:
```go
package main

import (
    "strconv"
    "time"
    "fmt"

    tc "github.com/nevadex/trackconvert"
)

func main() {
    sd := tc.AccurateSongData{Artist: "Pinkfong", Title: "Baby Shark"}
    sd, _ = time.ParseDuration("1m20s")
    song, err := tc.ConvertSongAccurate(sd)
    
    if err == nil {
    	fmt.Println(x.VideoId + " " + x.SearchResult.Channel + " " + x.SearchResult.Title + " " + strconv.Itoa(x.Accuracy))
    } else {
    	// handle error
    }

    result, err := tc.GetFirstResult("Pinkfong", "Baby Shark")
    if err == nil {
        fmt.Println(result.VideoId)
    } else {
    	// handle error
    } 

    results, err := tc.GetAllResults("Pinkfong", "Baby Shark")
    if err == nil {
    	for  _, i := range results {
    		fmt.Println(result.VideoId)
    	}
    } else {
    	// handle error
    }
}
```