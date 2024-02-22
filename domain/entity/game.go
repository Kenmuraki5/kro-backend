package entity

type Game struct {
	Id    string
	Name  string `json:"name"`
	Stock string `json:"stock"`
	Price string `json:"price"`
	Image string `json:"image"`
}
