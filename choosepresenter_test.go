package choosepresenter_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/jack-davidson/choosepresenter"
	"testing"
)

type FakeRand struct {
	Float64Func func() float64
}

func (r FakeRand) Float64() float64 {
	return r.Float64Func()
}

func TestChoosePresenter(t *testing.T) {
	got := choosepresenter.CalculateScores(
		FakeRand{Float64Func: func() float64 { return 0.2 }},
		[]choosepresenter.Participant{
			{"Jack", 1, 0},
			{"John", 0, 0},
			{"Isaac", 0, 0},
		})

	want := []choosepresenter.Participant{
		{"Jack", 1, 5.2},
		{"John", 0, 10.2},
		{"Isaac", 0, 10.2},
	}

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(got, want))
	}
}
