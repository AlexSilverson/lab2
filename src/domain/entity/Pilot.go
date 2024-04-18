package entity

type Pilot struct {
	ID         uint   `json:"id"`
	FirstName  string `json:"firstname"`
	SecondName string `json:"secondname"`
	LastName   string `json:"lastname"`
	Age        uint   `json:"age"`
	Experience uint   `json:"experience"`
}
