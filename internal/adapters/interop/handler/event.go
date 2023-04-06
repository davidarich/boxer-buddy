package handler

// Structure used for interop between Go and client
type InteropMessage struct {
	Type string   `json:"type"`
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}
