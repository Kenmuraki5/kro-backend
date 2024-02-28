package restmodel

type Game struct {
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Image       []string `json:"image"`
	ReleaseDate string   `json:"releaseDate"`
	SupDevice   []string `json:"supDevice"`
	Genre       []string `json:"genre"`
}
