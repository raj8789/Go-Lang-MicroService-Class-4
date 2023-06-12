package data

import (
	"fmt"
	"testing"
)

 func TestCheckValidation(t *testing.T){
		P:=&Product{
			Name:"Tea",
			Price:30,
			Description:"Tasty Tea",
			Sku :"abc-abc-abc",
		}
		err:=P.ValidateProduct()
		fmt.Println(err)
		if err!=nil {
			t.Fail()
		}else {
			t.Log("Test Success")
		}
 }