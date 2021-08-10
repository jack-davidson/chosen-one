package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jack-davidson/chosen-one/choose"
)

type SlackMessage struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	Blocks  string `json:"blocks"`
}

const Host = "localhost"
const Port = 8000
const ParticipantDataFileName = "participantdata.json"
const PostMessageURL = "https://slack.com/api/chat.postMessage"

func ReadToken(file string) string {
	token, err := os.ReadFile(file)
	if err != nil {
		return ""
	}
	strippedToken := strings.Replace(string(token), "\n", "", -1)
	return strippedToken
}

func PostMessage(msg SlackMessage, token string) error {
	payload, _ := json.Marshal(msg)

	req, err := http.NewRequest("POST", PostMessageURL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Failed to create request.")
		return errors.New("Failed to create request")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request.")
		return errors.New("Failed to send request.")
	}

	defer resp.Body.Close()
	return nil
}

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
	PostMessage(SlackMessage{
		Channel: r.FormValue("channel_id"),
		Blocks: `[
				{
					"type": "section",
					"text": {
						"type": "mrkdwn",
						"text": ">` + winner.Name + ` has been selected to present for this week!"
					}
				},
				{
					"type": "context",
					"elements": [
						{
							"type": "plain_text",
							"text": "/chooseperson ` + text + `",
							"emoji": true
						}
					]
				}
			]`,
	}, ReadToken("slack_token"))
}

func main() {
	http.HandleFunc("/chosen-one/commands/chooseperson", choosePerson)
	fmt.Printf("Listening on http://%s:%d\n", Host, Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", Host, Port), nil)
}
