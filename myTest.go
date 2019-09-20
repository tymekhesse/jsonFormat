package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
)




func main() {

	file, e := ioutil.ReadFile("./outputnew.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	myJson := string(file)
	m, ok := gjson.Parse(myJson).Value().(interface{})
	if !ok {
		fmt.Println("Error")
	}

	data :=	iterate(m)
	myJson2,_ := json.Marshal(&data)
	fmt.Println(string(myJson2))
}
