package restmodel

type Game struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Images      []string `json:"images"`
	SupDevice   []string `json:"supDevice"`
	Genre       []string `json:"genre"`
	Brand       string   `json:"brand"`
}
