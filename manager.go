package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Hello is a simple hello function to test working
func Hello() {
	fmt.Println("Hello manager")
}

// RunGame is the main method to manage the game life cycle between bots and game engine
func RunGame(engineURL string) {
	// gameEngine := make(chan string)
	// until game ends
	players := []Player{{PID: 0, Name: "Eric"}, {PID: 1, Name: "Sam"}}
	onCreate(engineURL, players)
	onStart(engineURL)
	for _, player := range players {
		getState(engineURL, player.PID)
	}
}

func onCreate(engineURL string, players []Player) {
	onCreateJSON, _ := json.Marshal(Action{
		Type:    "onCreate",
		Payload: players,
	})
	fmt.Println(string(onCreateJSON))
	httpPost(engineURL+"/game-engine", onCreateJSON)
}

func onStart(engineURL string) {
	onCreateJSON, _ := json.Marshal(Action{
		Type: "onStart",
	})
	fmt.Println(string(onCreateJSON))
	httpPost(engineURL+"/game-engine", onCreateJSON)
}

func getState(engineURL string, playerID int) {
	onCreateJSON, _ := json.Marshal(Action{
		Type:    "getState",
		Payload: Player{PID: playerID},
	})
	fmt.Println(string(onCreateJSON))
	httpPost(engineURL+"/game-engine", onCreateJSON)
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

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body)
}
