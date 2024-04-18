package entity

type Flight struct {
	ID         uint   `json:"id"`
	From       uint   `json:"from"`
	To         uint   `json:"to"`
	TrevelTime uint   `json:"treveltime"`
	Pilot      uint   `json:"pilot"`
	StartTime  string `json:"starttime"`
}
