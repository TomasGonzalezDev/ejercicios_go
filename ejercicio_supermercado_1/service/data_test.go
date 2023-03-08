package service

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestAllProducts(t *testing.T){

	client := &http.Client{}

	req , err1 := http.NewRequest("GET","http://localhost:8080/products",nil)
	fmt.Println(req,"tomas1")
	if err1 != nil {
		fmt.Errorf("error1")
	}

	response, _ := client.Do(req)

	fmt.Println(response,"tomas")

	if response == nil {
		fmt.Errorf("error 2")
	}

	defer response.Body.Close()

	body, err2 := io.ReadAll(response.Body)

	if err2 != nil {
		fmt.Errorf("---------------")
	}

	fmt.Println(body)




	//body, err2 := io.ReadAll(req.Request.Response.Body)
	
	fmt.Println(req)
	//assert.Equal(t,200,req.Request.Response.StatusCode)
	
}
