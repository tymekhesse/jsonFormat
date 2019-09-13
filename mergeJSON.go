package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type jsonObj []struct {
	Source struct {
		Layers struct {
			HTTPFileData []string `json:"http.file_data"`
		} `json:"layers"`
	} `json:"_source"`
}

func parseFile(filePath string) {
	fmt.Printf("// reading file %s\n", filePath)
	file, err1 := ioutil.ReadFile(filePath)
	if err1 != nil {
		fmt.Printf("// error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}
	var obj jsonObj

	err2 := json.Unmarshal(file, &obj)
	if err2 != nil {
		fmt.Println("error:", err2)
		os.Exit(1)
	}

	var merged interface{}
	var unmerged interface{}
	for i := range obj {
		str := obj[i].Source.Layers.HTTPFileData[0]
		if err := json.Unmarshal([]byte(str), &unmerged); err != nil {
			fmt.Println("error", err)
		}
		merged = merge(merged, unmerged)
	}
	if mergedjson, err := json.Marshal(merged); err == nil {
		ioutil.WriteFile("./merged.json", mergedjson, 0644)
	} else {
		fmt.Println("error", err)
	}
}

func merge(merged interface{}, unmerged interface{}) interface{} {
	if merged == nil {
		return unmerged
	}
	unmergedmap, ok := unmerged.(map[string]interface{})
	if !ok {
		// if not a map then we can skip
		return nil
	}
	mergedmap := merged.(map[string]interface{})
	for k, v := range unmergedmap {
		if _, ok := mergedmap[k]; ok {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice:
				for _, item := range v.([]interface{}) {
					fmt.Println(mergedmap[k].([]interface{})[0], item)
					merge(mergedmap[k].([]interface{})[0], item)
				}
			case reflect.Map:
				merge(mergedmap[k], v)
			}
		} else {
			mergedmap[k] = v
		}
	}
	return interface{}(mergedmap)
}
