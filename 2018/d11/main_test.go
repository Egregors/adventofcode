package main

import (
	"reflect"
	"testing"
)

func TestRack_getPowerLvl(t *testing.T) {
	type fields struct {
		serial int
		x      int
		y      int
		ID     int
		power  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"8 3,5 = 4",
			fields{serial: 8, x: 3, y: 5},
			4,
		},
		{
			"57 122,79 = -5",
			fields{serial: 57, x: 122, y: 79},
			-5,
		},
		{
			"39 217,196 = 0",
			fields{serial: 39, x: 217, y: 196},
			0,
		},
		{
			"71 101,153 = 4",
			fields{serial: 71, x: 101, y: 153},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRack(tt.fields.serial, tt.fields.x, tt.fields.y)
			if got := r.getPowerLvl(); got != tt.want {
				t.Errorf("getPowerLvl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_getBestTotalPower(t *testing.T) {
	tests := []struct {
		name      string
		serial    int
		wantX     int
		wantY     int
		wantPower int
	}{
		{"18", 18, 33, 45, 29},
		{"42", 42, 21, 61, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid(tt.serial)
			gotX, gotY := g.getBestTotalPower()
			if gotX != tt.wantX {
				t.Errorf("getBestTotalPower() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("getBestTotalPower() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestNewGrid(t *testing.T) {
	type args struct {
		serial int
	}
	tests := []struct {
		name string
		args args
		want *Grid
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrid(tt.args.serial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRack(t *testing.T) {
	type args struct {
		serial int
		x      int
		y      int
	}
	tests := []struct {
		name string
		args args
		want *Rack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRack(tt.args.serial, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRack_genID(t *testing.T) {
	type fields struct {
		serial int
		x      int
		y      int
		ID     int
		power  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rack{
				serial: tt.fields.serial,
				x:      tt.fields.x,
				y:      tt.fields.y,
				ID:     tt.fields.ID,
				power:  tt.fields.power,
			}
			if got := r.genID(); got != tt.want {
				t.Errorf("genID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRack_getPowerLvl1(t *testing.T) {
	type fields struct {
		serial int
		x      int
		y      int
		ID     int
		power  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rack{
				serial: tt.fields.serial,
				x:      tt.fields.x,
				y:      tt.fields.y,
				ID:     tt.fields.ID,
				power:  tt.fields.power,
			}
			if got := r.getPowerLvl(); got != tt.want {
				t.Errorf("getPowerLvl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHundredsDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHundredsDigit(tt.args.n); got != tt.want {
				t.Errorf("getHundredsDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
