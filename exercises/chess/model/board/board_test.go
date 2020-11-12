package board

import (
	"reflect"
	"testing"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/player"
)

type mockCoord int
type mockPiece struct {
}

func (p mockPiece) String() string {
	panic("implement me")
}

func (p mockPiece) Moves(isCapture bool) map[coord.ChessCoordinates]bool {
	panic("implement me")
}

func (p mockPiece) Color() player.Color {
	panic("implement me")
}

// Coord returns x if n==0, y if n==1
func (c mockCoord) Coord(n int) (int, error) {
	return int(c), nil
}

func (c mockCoord) String() string {
	return "1"
}

func TestClassic_MovePiece(t *testing.T) {
	pi := mockPiece{}
	pi1 := mockPiece{}
	class := Classic{}
	coordinates := coord.NewCartesian(0, 0)
	x, _ := coordinates.Coord(0)
	y, _ := coordinates.Coord(1)
	class[x][y] = pi
	coordinates1 := coord.NewCartesian(3, 5)
	x1, _ := coordinates1.Coord(0)
	y1, _ := coordinates1.Coord(1)
	class[x1][y1] = pi1

	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool
	}{
		{
			"Successful move",
			class,
			args{from: coordinates, to: coord.NewCartesian(4, 5)},
			false,
		},
		{
			"No piece found at from",
			class,
			args{from: coord.NewCartesian(7, 0), to: coordinates},
			true,
		},
		{
			"Occupied at to",
			class,
			args{from: coordinates1, to: coordinates},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.MovePiece(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Classic.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClassic_PlacePieceAt(t *testing.T) {
	pi := mockPiece{}
	class := Classic{}
	coordinates := coord.NewCartesian(0, 0)
	x, _ := coordinates.Coord(0)
	y, _ := coordinates.Coord(1)
	class[x][y] = pi
	type args struct {
		p  piece.Piece
		at coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool
	}{
		{
			"Failed",
			class,
			args{p: pi, at: coordinates},
			true,
		},
		{
			"Passed",
			class,
			args{p: pi, at: coord.NewCartesian(2, 5)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.PlacePieceAt(tt.args.p, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("Classic.PlacePieceAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClassic_PieceAt(t *testing.T) {
	pi := mockPiece{}
	class := Classic{}
	coordinates := coord.NewCartesian(0, 0)
	x, _ := coordinates.Coord(0)
	y, _ := coordinates.Coord(1)
	class[x][y] = pi

	type args struct {
		at coord.ChessCoordinates
	}
	tests := []struct {
		name string
		c    Classic
		args args
		want piece.Piece
	}{{
		"A1",
		class,
		args{at: coordinates},
		pi,
	}, {
		"C6",
		class,
		args{at: coord.NewCartesian(2, 5)},
		nil,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.PieceAt(tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Classic.PieceAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
