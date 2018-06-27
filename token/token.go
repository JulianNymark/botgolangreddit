package token

import (
    "io/ioutil"
    "fmt"
    
)

var fileName string = "./token/jwt"

func Get()(token string, err error){
	temp , err := ioutil.ReadFile(fileName)
	if err != nil {
	} else {
	token = string(temp)
	}
	return 
}

func Set(token string)(err error){
	b := []byte(token)
	err = ioutil.WriteFile(fileName, b, 0644)
	return
}