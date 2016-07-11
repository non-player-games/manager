package manager

// Action is DTO (data transfer object) between game engine and manager
// note that payload is a dynamic type (thus interface{})
type Action struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// Player struct to indicate each player differently
type Player struct {
	PID  int    `json:"pid"`
	Name string `json:"name"`
}

// Move represent single move from single player
type Move struct {
	PID  int         `json:"pid"`
	Move interface{} `json:"move"`
}

// Moves is a DTO to wrap moves and game id together
type Moves struct {
	GID   int    `json:"gid"`
	Moves []Move `json:"moves"`
}
