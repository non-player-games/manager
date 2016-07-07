package manager

import (
	"encoding/json"
	"fmt"
)

// A simple hello function to test working
func Hello() {
	fmt.Println("Hello manager")
}

// Action is DTO (data transfer object) between game engine and manager
type Action struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

// Player struct to indicate each player differently
type Player struct {
	PID  string `json:"pid"`
	Name string `json:"name"`
}

// RunGame is the main method to manage the game life cycle between bots and game engine
func RunGame(engineURL string) {
	// gameEngine := make(chan string)
	// until game ends
	for {
		onCreate := &Action{
			Type:    "onCreate",
			Payload: "",
		}
		onCreateJSON, _ := json.Marshal(onCreate)
		fmt.Println(string(onCreateJSON))
	}
}
