package main

import "testing"

func Test_playGame(t *testing.T) {
	type args struct {
		playersNumber int
		lastMarble    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"9  – 25   - 32", args{9, 25}, 32},
		{"10 – 1618 - 8317", args{10, 1618}, 8317},
		{"13 – 7999 - 146373", args{13, 7999}, 146373},
		{"17 – 1104 - 2764", args{17, 1104}, 2764},
		{"21 – 6111 - 54718", args{21, 6111}, 54718},
		{"30 – 5807 - 37305", args{30, 5807}, 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playGame(tt.args.playersNumber, tt.args.lastMarble); got != tt.want {
				t.Errorf("playGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
