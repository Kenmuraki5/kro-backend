package restmodel

type Console struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Images      []string `json:"images"`
}
