package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type jsonObj []struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	Score  interface{} `json:"_score"`
	Source struct {
		Layers struct {
			HTTPFileData []string `json:"http.file_data"`
		} `json:"layers"`
	} `json:"_source"`
}

func main() {
	rawFileInputPath := "./output3.json"
	slice := make([]string, 0)
	file := readFile(rawFileInputPath)
	jsonFromFile := fileToJson(file)
	slice = jsonToSlice(jsonFromFile)


	writeToFile(slice)
}

func readFile(filePath string) (file []byte) {

	fmt.Printf("// reading file %s\n", filePath)
	file, err1 := ioutil.ReadFile(filePath)
	if err1 != nil {
		fmt.Printf("// error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}
	return
}
func jsonToSlice(obj jsonObj) (slice []string) {

	for i, _ := range obj {
		str := obj[i].Source.Layers.HTTPFileData[0]
		slice = append(slice, str)
	}
	return
}

func fileToJson(file []byte) (obj jsonObj) {
	err2 := json.Unmarshal(file, &obj)
	if err2 != nil {
		fmt.Println("error:", err2)
		os.Exit(1)
	}
	return
}
func writeToFile(slice []string, isSchema bool) {

	f, err := os.Create("output3formatted.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var sb strings.Builder
	if !isSchema {
		sb.WriteString("[")
	}
	for i := 0; i < len(slice); i++ {
		sb.WriteString(slice[i])
		if i < len(slice)-1 {
			if isSchema {
				sb.WriteString("\r\n")
			} else {
				sb.WriteString(",")
			}
		}
	}
	if !isSchema {
		sb.WriteString("]")
	}
	l, err := f.WriteString(sb.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
