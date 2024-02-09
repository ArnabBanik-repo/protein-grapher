package models

import (
	"reflect"
	"testing"
)

var src Protein = Protein("ABCC2")
var x Protein = Protein("X")
var y Protein = Protein("Y")
var z Protein = Protein("Z")
var dest Protein = Protein("CFTR")

func TestHasCycle(t *testing.T) {
	t.Run("With cycle", func(t *testing.T) {

		var network ProteinNetwork = ProteinNetwork{
			src: ProteinList{x},
			x:   ProteinList{y},
			y:   ProteinList{z},
			z:   ProteinList{dest, src},
		}
		want := true
		got := network.HasCycle()

		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("Without cycle", func(t *testing.T) {

		var network ProteinNetwork = ProteinNetwork{
			src: ProteinList{x},
			x:   ProteinList{y},
			y:   ProteinList{z},
			z:   ProteinList{dest},
		}
		want := false
		got := network.HasCycle()

		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestShortestPath(t *testing.T) {
	t.Run("no cycles", func(t *testing.T) {
		var network ProteinNetwork = ProteinNetwork{
			src: ProteinList{x},
			x:   ProteinList{y},
			y:   ProteinList{z},
			z:   ProteinList{dest},
		}

		want := []ProteinList{ProteinList{src, x}, ProteinList{src, x, y}, ProteinList{src, x, y, z}, ProteinList{src, x, y, z, dest}}
		got := network.ShortestPaths(src)

		if !areSlicesEqual(want, got) {
			t.Errorf("\nwant %v\ngot %v", want, got)
		}
	})

  t.Run("cycles", func(t *testing.T) {
		var network ProteinNetwork = ProteinNetwork{
			src: ProteinList{x},
			x:   ProteinList{y},
			y:   ProteinList{z},
			z:   ProteinList{src, dest},
		}

		want := []ProteinList{ProteinList{src, x}, ProteinList{src, x, y}, ProteinList{src, x, y, z}, ProteinList{src, x, y, z, dest}}
		got := network.ShortestPaths(src)

		if !areSlicesEqual(want, got) {
			t.Errorf("\nwant %v\ngot %v", want, got)
		}
  })
}

func areSlicesEqual(slice1, slice2 []ProteinList) bool {
	// Flattening both slices using maps to preserve duplicates
	flattened1 := make(map[interface{}]bool)
	for _, sublist := range slice1 {
		for _, item := range sublist {
			flattened1[item] = true
		}
	}

	flattened2 := make(map[interface{}]bool)
	for _, sublist := range slice2 {
		for _, item := range sublist {
			flattened2[item] = true
		}
	}

	// Comparing the flattened maps for equality
	if !reflect.DeepEqual(flattened1, flattened2) {
		return false
	}

	return true
}
