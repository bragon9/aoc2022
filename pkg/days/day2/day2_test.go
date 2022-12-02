package day2

import (
	"testing"
)

func Test_chooseMove(t *testing.T) {
	type args struct {
		enemyMove move
		outcome   outcome
	}
	tests := []struct {
		name string
		args args
		want move
	}{
		{
			name: "Win vs Rock",
			args: args{
				enemyMove: Rock,
				outcome:   Win,
			},
			want: Paper,
		},
		{
			name: "Loss vs Rock",
			args: args{
				enemyMove: Rock,
				outcome:   Loss,
			},
			want: Scissors,
		},
		{
			name: "Draw vs Rock",
			args: args{
				enemyMove: Rock,
				outcome:   Draw,
			},
			want: Rock,
		},
		{
			name: "Win vs Paper",
			args: args{
				enemyMove: Paper,
				outcome:   Win,
			},
			want: Scissors,
		},
		{
			name: "Loss vs Paper",
			args: args{
				enemyMove: Paper,
				outcome:   Loss,
			},
			want: Rock,
		},
		{
			name: "Draw vs Paper",
			args: args{
				enemyMove: Paper,
				outcome:   Draw,
			},
			want: Paper,
		},
		{
			name: "Win vs Scissors",
			args: args{
				enemyMove: Scissors,
				outcome:   Win,
			},
			want: Rock,
		},
		{
			name: "Loss vs Scissors",
			args: args{
				enemyMove: Scissors,
				outcome:   Loss,
			},
			want: Paper,
		},
		{
			name: "Draw vs Scissors",
			args: args{
				enemyMove: Scissors,
				outcome:   Draw,
			},
			want: Scissors,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chooseMove(tt.args.enemyMove, tt.args.outcome); got != tt.want {
				t.Errorf("chooseMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOutcome(t *testing.T) {
	type args struct {
		enemyMove move
		myMove    move
	}
	tests := []struct {
		name string
		args args
		want outcome
	}{
		{
			name: "My Rock vs Enemy Rock",
			args: args{
				enemyMove: Rock,
				myMove:    Rock,
			},
			want: Draw,
		},
		{
			name: "My Rock vs Enemy Scissors",
			args: args{
				enemyMove: Scissors,
				myMove:    Rock,
			},
			want: Win,
		},
		{
			name: "My Rock vs Enemy Paper",
			args: args{
				enemyMove: Paper,
				myMove:    Rock,
			},
			want: Loss,
		},
		{
			name: "My Paper vs Enemy Rock",
			args: args{
				enemyMove: Rock,
				myMove:    Paper,
			},
			want: Win,
		},
		{
			name: "My Paper vs Enemy Scissors",
			args: args{
				enemyMove: Scissors,
				myMove:    Paper,
			},
			want: Loss,
		},
		{
			name: "My Paper vs Enemy Paper",
			args: args{
				enemyMove: Paper,
				myMove:    Paper,
			},
			want: Draw,
		},
		{
			name: "My Scissors vs Enemy Rock",
			args: args{
				enemyMove: Rock,
				myMove:    Scissors,
			},
			want: Loss,
		},
		{
			name: "My Scissors vs Enemy Scissors",
			args: args{
				enemyMove: Scissors,
				myMove:    Scissors,
			},
			want: Draw,
		},
		{
			name: "My Scissors vs Enemy Paper",
			args: args{
				enemyMove: Paper,
				myMove:    Scissors,
			},
			want: Win,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOutcome(tt.args.enemyMove, tt.args.myMove); got != tt.want {
				t.Errorf("getOutcome() = %v, want %v", got, tt.want)
			}
		})
	}
}
