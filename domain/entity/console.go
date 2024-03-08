package entity

type Console struct {
	Id          string   `json:"Id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Images      []string `json:"images"`
	ReleaseDate string   `json:"releaseDate"`
}
