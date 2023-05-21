package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Sku         string  `json:"-"`
	Create      string  `json:"-"`
	Update      string  `json:"-"`
	Deletion    string  `json:"-"`
}

var Productlist = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Chill and Relaxed Coffee",
		Price:       100.0,
		Sku:         "abc234",
		Create:      time.Now().UTC().String(),
		Update:      time.Now().UTC().String(),
		Deletion:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Musk",
		Description: "Cold and Relaxed Coffee",
		Price:       100.0,
		Sku:         "abc234",
		Create:      time.Now().UTC().String(),
		Update:      time.Now().UTC().String(),
		Deletion:    time.Now().UTC().String(),
	},
}

func GetProductList() []*Product {
	return Productlist
}
