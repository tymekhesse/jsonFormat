package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	parseFile("./output3.json")

	file, e := ioutil.ReadFile("./merged.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var v interface{}
	if err := json.Unmarshal(file, &v); err != nil {
		fmt.Println("error", err)
	}

	data := iterate(v)
	schemaJSON, _ := json.MarshalIndent(&data, "", "  ")
	ioutil.WriteFile("schema.json", schemaJSON, 0644)
}
