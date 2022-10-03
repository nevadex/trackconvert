package trackconvert

// import (
// 	"strings"
// 	"unicode"
// 	"github.com/buger/jsonparser"
// )

func TrackFromString(inputName string) Track {
	inputName = cleanString(inputName)

	return Track{id: "nil", accuracy: 0}
}
