package chosenone_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/jack-davidson/chosenone"
	"testing"
)

type FakeRand struct {
	Float64Func func() float64
}

func (r FakeRand) Float64() float64 {
	return r.Float64Func()
}

func TestCalculateScores(t *testing.T) {
	got := chosenone.CalculateScores(
		FakeRand{Float64Func: func() float64 { return 0.2 }},
		[]chosenone.Participant{
			{"Jack", 1, 0},
			{"John", 0, 0},
			{"Isaac", 0, 0},
		})

	want := []chosenone.Participant{
		{"Jack", 1, 5.2},
		{"John", 0, 10.2},
		{"Isaac", 0, 10.2},
	}

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(got, want))
	}
}

func TestWinner(t *testing.T) {
	want := chosenone.Participant{"John", 0, 10.9}
	got := chosenone.Winner([]chosenone.Participant{
		{"Jack", 1, 5.2},
		{"Isaac", 0, 10.2},
		{"John", 0, 10.9},
	})
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(got, want))
	}
}
