package main_test

import (
	"io/ioutil"
	"testing"

	"github.com/go-test/deep"
)

type T struct {
	Name string

	Numbers []float64
}

func TestMergeGpx(t *testing.T) {
	// TODO: create mocks and use testify
	// GIVEN
	// mockRoutes

	// WHEN
	// call MergeGPX with those routes

	// THEN
	// expect the mock and the generated to be equal
	generated, _ := ioutil.ReadFile("./merged.gpx")
	mock, _ := ioutil.ReadFile("./mockMerged.gpx")

	if diff := deep.Equal(generated, mock); diff != nil {
		t.Error(diff)
	}
}
