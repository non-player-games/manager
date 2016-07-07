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
	onCreateRequest(engineURL)
}

func onCreateRequest(engineURL string) {
	playerJSON, _ := json.Marshal([2]Player{{PID: 0, Name: "Eric"}, {PID: 1, Name: "Sam"}})
	onCreateJSON, _ := json.Marshal(Action{
		Type:    "onCreate",
		Payload: string(playerJSON),
	})
	httpPost(engineURL+"/game-engine", onCreateJSON)

	fmt.Println(string(onCreateJSON))
}

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
