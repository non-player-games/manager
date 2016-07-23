package manager

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple hello function to test working
func Hello() {
	log.Println("Hello manager")
}

// RunGame is the main method to manage the game life cycle between bots and game engine
func RunGame(engineURL string) {
	// gameEngine := make(chan string)
	// until game ends
	players := []Player{{PID: 0, Name: "Eric"}}
	onCreate(engineURL, players)
	onStart(engineURL)
	var botMove string
	for _, player := range players {
		playerState := getState(engineURL, player.PID)
		botMove = botTurnStart(playerState)
	}
	result := processMoves(engineURL, botMove)
	log.Println(result)
}

func onCreate(engineURL string, players []Player) {
	onCreateJSON, _ := json.Marshal(Action{
		Type:    "onCreate",
		Payload: players,
	})
	log.Println(string(onCreateJSON))
	httpPost(engineURL+"/game-engine", onCreateJSON)
}

func onStart(engineURL string) {
	onCreateJSON, _ := json.Marshal(Action{
		Type: "onStart",
	})
	log.Println(string(onCreateJSON))
	httpPost(engineURL+"/game-engine", onCreateJSON)
}

func getState(engineURL string, playerID int) string {
	onCreateJSON, _ := json.Marshal(Action{
		Type:    "getState",
		Payload: Player{PID: playerID},
	})
	log.Println(string(onCreateJSON))
	return httpPost(engineURL+"/game-engine", onCreateJSON)
}

// as for now, bot turn start is random move for our game engine (bot saving
// prince).
// TODO: think about how to pass bot into this argument
func botTurnStart(botState string) string {
	return "n"
}

func processMoves(engineURL string, botMove string) string {
	processMoveJSON, _ := json.Marshal(Action{
		Type: "processMove",
		Payload: Moves{
			GID: 0,
			Moves: []Move{
				Move{
					PID:  0,
					Move: botMove,
				},
			},
		},
	})
	return httpPost(engineURL, processMoveJSON)
}

// abstract out the http call boiler plate and error handling
func httpPost(url string, payload []byte) string {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return string(body)
}
