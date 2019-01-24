package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSort(t *testing.T) {
	sortData := []servers{
		{"Singapore #23", 1023},
		{"Germany #36", 1979},
		{"Lithuania #15", 800},
		{"United Kingdom #58", 388},
	}
	sortExpectedBest := []servers{
		{"United Kingdom #58", 388},
		{"Lithuania #15", 800},
		{"Singapore #23", 1023},
		{"Germany #36", 1979},
	}
	sortExpectedAbc := []servers{
		{"Germany #36", 1979},
		{"Lithuania #15", 800},
		{"Singapore #23", 1023},
		{"United Kingdom #58", 388},
	}
	Sort(sortData, "best")
	if !cmp.Equal(sortData, sortExpectedBest) {
		t.Errorf("Sorting servers by best failed. Expected: %v. Got %v", sortExpectedBest, sortData)
	}
	Sort(sortData, "alphabetical")
	if !cmp.Equal(sortData, sortExpectedAbc) {
		t.Errorf("Sorting servers by best failed. Expected: %v. Got %v", sortExpectedAbc, sortData)
	}
}
