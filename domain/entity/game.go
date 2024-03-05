package entity

type Game struct {
	Id          string   `json:"Id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Image       []string `json:"image"`
	ReleaseDate string   `json:"releaseDate"`
	SupDevice   []string `json:"supDevice"`
	Genre       []string `json:"genre"`
	Brand       string   `json:"brand"`
}
