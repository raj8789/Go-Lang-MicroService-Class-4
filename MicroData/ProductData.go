package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

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

type Products []*Product

func (pr *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(pr)
}

func GetProductList2() Products {
	return Productlist
}
func (pr *Product) FromJson(b []byte) error {
	return json.Unmarshal(b, &pr)
}
func (pr *Product) AddProductToList() {
	id := getNextId()
	pr.ID = id
	Productlist = append(Productlist, pr)
}
func getNextId() int {
	pl := Productlist[len(Productlist)-1]
	return pl.ID + 1
}

var ErrProductnotFound = fmt.Errorf("Product Not Found")

func UpdateProduct(id int, prod *Product) error {
	fp, fid, e := finProduct(id)
	if e != nil {
		return e
	}
	fp.ID = id
	prod.ID = id
	Productlist[fid] = prod
	return nil
}
func finProduct(id int) (*Product, int, error) {
	for i, pr := range Productlist {
		if pr.ID == id {
			return pr, i, nil
		}
	}
	return nil, -1, ErrProductnotFound
}
