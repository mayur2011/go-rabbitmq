package model

type Message struct {
	Sequence int    `json:"seq"`
	Payload  string `json:"payload"`
}
