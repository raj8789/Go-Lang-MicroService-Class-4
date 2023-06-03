package handler

import (
	data "MicroData"
	"encoding/json"
	"fmt"

	"io/ioutil"

	//"io"
	"log"
	"net/http"
)

type ProductHandlerList struct {
	l    *log.Logger
	name string
}

func NewProductHandlerList(l *log.Logger, name string) *ProductHandlerList {
	return &ProductHandlerList{l, name}
}

func (productHandlerList *ProductHandlerList) ServeHTTP(rw http.ResponseWriter, re *http.Request) {

	//********************* Way 1 for get request in go lang******************************
	if re.Method == http.MethodGet {
		pl := data.GetProductList()
		d, err := json.Marshal(pl)
		if err != nil {
			http.Error(rw, "Unable to Parse Json Data", http.StatusBadRequest)
		} else {
			fmt.Printf("% s Your Code Runs Fine", productHandlerList.name)
			rw.Write(d)
		}
	}

	//********************* Way 2 for get request in go lang******************************
	if re.Method == http.MethodGet {
		productHandlerList.getProducts(rw, re)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)

	//**************************** post request to add a Product in Productlist *****************************
	if re.Method == http.MethodPost {
		productHandlerList.addProduct(rw, re)
	}
}
func (productHandlerList *ProductHandlerList) getProducts(rw http.ResponseWriter, re *http.Request) {
	listproduct := data.GetProductList2()
	err := listproduct.ToJson(rw)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
	}
}
func (productHandlerList *ProductHandlerList) addProduct(rw http.ResponseWriter, re *http.Request) {
	productHandlerList.l.Println("Http Post Method to Add Product To ProductList")
	prod := data.Product{}
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "Error reading request body:", err)
		return
	}
	fmt.Printf("Hiii Body Of Request=%s", string(body))
	err = prod.FromJson(body)
	fmt.Println("Hiii Error Message=%s", err)
	if err != nil {
		http.Error(rw, "Unable to Decode Json to Go Value", http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Request processed successfully"))
	}
	productHandlerList.l.Printf("Product=#%v", prod)
	prod.AddProductToList()
}
