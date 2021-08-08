package choosepresenter

import ()

type Randomer interface {
	Float64() float64
}

type Participant struct {
	Name              string
	PresentationCount int
	Score             float64
}

func ChoosePresenter(randomer Randomer, participants []Participant) []Participant {
	for _, participant := range participants {
		participant.Score += (1 / (1 + float64(participant.PresentationCount))) + randomer.Float64()
	}
	return participants
}
