package main

import "testing"

func TestGrid_getBestTotalPower(t *testing.T) {

	tests := []struct {
		name     string
		serial   int
		wantX    int
		wantY    int
		wantSize int
	}{
		{"18", 18, 90, 269, 16},
		{"42", 42, 232, 251, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid(tt.serial)
			gotX, gotY, gotSize := g.getBestTotalPower()
			if gotX != tt.wantX {
				t.Errorf("getBestTotalPower() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("getBestTotalPower() gotY = %v, want %v", gotY, tt.wantY)
			}
			if gotSize != tt.wantSize {
				t.Errorf("getBestTotalPower() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}
