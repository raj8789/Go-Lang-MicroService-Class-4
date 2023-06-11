package handler

import (
	data "MicroData"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

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

	//**************************** put request to update a Product in productlist ***************************
	if re.Method == http.MethodPut {
		r := regexp.MustCompile(`/([0-9]+)`)
		g := r.FindAllStringSubmatch(re.URL.Path, -1)
		fmt.Println("Value of g=%s", g)
		if len(g) != 1 {
			http.Error(rw, "Invalid Url", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid Url", http.StatusBadRequest)
			return
		}
		fmt.Println("Value of g[0]=%s", g[0])
		idString := g[0][1]
		fmt.Println("Value of g[0][0]=%s", g[0][0])
		fmt.Println("Value of g[0][1]=%s", g[0][1])
		id, _ := strconv.Atoi(idString)
		productHandlerList.l.Println("Id REceived", id)
		productHandlerList.updateProduct(id, rw, re)
	}
}
func (ProductHandlerList *ProductHandlerList) updateProduct(id int, rw http.ResponseWriter, re *http.Request) {
	prod := data.Product{}
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "Error reading request body:", err)
		return
	}
	fmt.Printf("Hiii Body Of Request=%s", string(body))
	err = prod.FromJson(body)
	if err != nil {
		http.Error(rw, "Unable to Decode Json to Go Value", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, &prod)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "Product Not Found To Update:", err)
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
