package ncaa

type Schedule struct {
	UpdatedAt string        `json:"updated_at"`
	Games     []GameWrapper `json:"games"`
}

type GameWrapper struct {
	Game GameNCAA `json:"game"`
}
