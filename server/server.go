package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/jack-davidson/chosen-one/choose"
)

const Host = "localhost"
const Port = 8000

func choosePerson(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	names := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' '
	})
	participants := make([]choose.Participant, len(names))
	for i, name := range names {
		participants[i].Name = name
	}
	winner := choose.Winner(choose.CalculateScores(rand.New(rand.NewSource(time.Now().UnixNano())), participants))
	fmt.Println(winner)
	fmt.Println(text)
	fmt.Println(names)
}

func main() {
	http.HandleFunc("/chosen-one/commands/chooseperson", choosePerson)
	fmt.Printf("Listening on http://%s:%d\n", Host, Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", Host, Port), nil)
}
