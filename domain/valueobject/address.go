package valueobject

type Address struct {
	Address     string `json:"address"`
	Province    string `json:"province"`
	District    string `json:"district"`
	SubDistrict string `json:"subDistrict"`
	PostalCode  string `json:"postalCode"`
}
