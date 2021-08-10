package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jack-davidson/chosen-one/choose"
)

const Host = "localhost"
const Port = 8000
const ParticipantDataFileName = "participantdata.json"

func choosePerson(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	names := strings.Fields(text)
	file, err := os.ReadFile(ParticipantDataFileName)
	if err != nil {
		fmt.Println(err)
	}
	participantData := make(map[string]interface{})
	json.Unmarshal(file, &participantData)
	participants := make([]choose.Participant, len(names))
	for i, name := range names {
		if participantData[name] != nil {
			participants[i].PresentationCount = int(participantData[name].(float64))
		}
		participants[i].Name = name
	}

	winner := choose.Winner(choose.CalculateScores(rand.New(rand.NewSource(time.Now().UnixNano())), participants))
	participantData[winner.Name] = winner.PresentationCount + 1
	file, err = json.Marshal(participantData)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile(ParticipantDataFileName, file, 0644)
}

func main() {
	http.HandleFunc("/chosen-one/commands/chooseperson", choosePerson)
	fmt.Printf("Listening on http://%s:%d\n", Host, Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", Host, Port), nil)
}
