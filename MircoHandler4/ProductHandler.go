package handler

import (
	data "MicroData"
	"encoding/json"
	"fmt"
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
	pl := data.GetProductList()
	d, err := json.Marshal(pl)
	if err != nil {
		http.Error(rw, "Unable to Parse Json Data", http.StatusBadRequest)
	} else {
		fmt.Printf("% s Your Code Runs Fine", productHandlerList.name)
		rw.Write(d)
	}

	//********************* Way 2 for get request in go lang******************************
	if re.Method == http.MethodGet {
		productHandlerList.getProducts(rw, re)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (productHandlerList *ProductHandlerList) getProducts(rw http.ResponseWriter, re *http.Request) {
	listproduct := data.GetProductList2()
	err := listproduct.ToJson(rw)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
	}
}

