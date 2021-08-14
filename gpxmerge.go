package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tkrajina/gpxgo/gpx"
)

// MergeGpx merges gpx files with a descending date order
func MergeGpx(path []string, gpxFiles []*gpx.GPX) []byte {
	gpxMerged := gpx.GPX{}

	if gpxFiles[0].TimeBounds().StartTime.Before(gpxFiles[1].TimeBounds().StartTime) {
		gpxFiles[0].AppendTrack(&gpxFiles[1].Tracks[0])
		gpxMerged = *gpxFiles[0]
	} else {
		gpxFiles[1].AppendTrack(&gpxFiles[0].Tracks[0])
		gpxMerged = *gpxFiles[1]
	}

	xmlBytes, err := gpxMerged.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})
	check(err)
	return xmlBytes
}

func readAndParse(path string) *gpx.GPX {
	gpxBytes, err := ioutil.ReadFile(path)
	check(err)
	gpxFile, err := gpx.ParseBytes(gpxBytes)
	check(err)
	return gpxFile
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// TODO: option to merge all *.gpx files from dir
	paths := []string{}
	gpxFiles := []*gpx.GPX{}

	if len(os.Args) == 1 {
		paths = []string{"./route1.gpx", "./route2.gpx"}
	} else {
		for i := 1; i < len(os.Args); i++ {
			paths = append(paths, os.Args[i])
		}
	}

	for i := 0; i < len(paths); i++ {
		gpxFiles = append(gpxFiles, readAndParse(paths[i]))
	}
	mergedXmlBytes := MergeGpx(paths, gpxFiles)

	f, err := os.Create("./merged.gpx")
	check(err)
	defer f.Close()

	n, err := f.Write(mergedXmlBytes)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
}
