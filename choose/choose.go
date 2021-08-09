package choose

import ()

type Randomer interface {
	Float64() float64
}

type Participant struct {
	Name              string
	PresentationCount int
	Score             float64
}

// Calculate random score.
// The score is calculated by taking a random number and adding it
// to ten over the previous presentation count plus one:
//	score = random() + (10/presentation_count+1)
func CalculateScores(randomer Randomer,
	participants []Participant) []Participant {
	for i := range participants {
		participants[i].Score += randomer.Float64() +
			(10 / float64(participants[i].PresentationCount+1))
	}
	return participants
}

// Select the winner (largest score) from slice of (Participant)s.
func Winner(participants []Participant) Participant {
	var chosenParticipant Participant
	for _, participant := range participants {
		if participant.Score > chosenParticipant.Score {
			chosenParticipant = participant
		}
	}
	return chosenParticipant
}
