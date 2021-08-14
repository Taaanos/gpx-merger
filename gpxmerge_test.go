package main

import (
	"io/ioutil"
	"testing"

	"github.com/go-test/deep"
	"github.com/tkrajina/gpxgo/gpx"
)

func TestMergeGpx(t *testing.T) {
	path1 := "./mocks/mockRoute1.gpx"
	path2 := "./mocks/mockRoute2.gpx"

	// GIVEN
	mMerged, _ := ioutil.ReadFile("./mocks/mockMerged.gpx")

	// WHEN
	merged := MergeGpx([]string{path1, path2}, []*gpx.GPX{readAndParse(path1), readAndParse(path2)})

	// THEN
	if diff := deep.Equal(merged, mMerged); diff != nil {
		t.Error(diff)
	}
}
