package chosenone

import ()

type Randomer interface {
	Float64() float64
}

type Participant struct {
	Name              string
	PresentationCount int
	Score             float64
}

func CalculateScores(randomer Randomer,
	participants []Participant) []Participant {
	for i := range participants {
		participants[i].Score += randomer.Float64() +
			(10 / float64(participants[i].PresentationCount+1))
	}
	return participants
}

func Winner(participants []Participant) Participant {
	var chosenParticipant Participant
	for _, participant := range participants {
		if participant.Score > chosenParticipant.Score {
			chosenParticipant = participant
		}
	}
	return chosenParticipant
}
