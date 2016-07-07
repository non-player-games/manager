package manager

// Action is DTO (data transfer object) between game engine and manager
type Action struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

// Player struct to indicate each player differently
type Player struct {
	PID  int    `json:"pid"`
	Name string `json:"name"`
}
