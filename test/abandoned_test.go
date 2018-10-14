package test

import (
	"testing"
	"github.com/a-h/generate/test/abandoned_gen"
)

func TestAbandoned(t *testing.T) {
	// this just tests the name generation works correctly
	r := abandoned.Root{
		Name: "jonson",
		Abandoned: &abandoned.PackageList{},
	}
	if r.Abandoned != nil {
		t.Fatal("thats the test")
	}
}
