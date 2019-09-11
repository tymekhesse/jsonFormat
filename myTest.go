package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// parseFile()

	file, e := ioutil.ReadFile("./outputnew.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var v interface{}
	if err := json.Unmarshal(file, &v); err != nil {
		fmt.Println("error", err)
	}

	data := iterate(v)
	myJSON2, _ := json.Marshal(&data)
	fmt.Println(string(myJSON2))
}
