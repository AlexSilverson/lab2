package entity

type City struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	History uint   `json:"history"`
	Resort  uint   `json:"resort"`
}
