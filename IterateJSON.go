package main

import (
	"reflect"
)

type schema struct {
	Name   string `json:"name"`
	Type  string	`json:"type"`
	Mode   string	`json:"mode"`
	Fields *[]schema `json:"fields,omitempty"`
}

func iterate(data interface{}) []schema {

	tmp := make([]schema, 0)


	if reflect.ValueOf(data).Kind() == reflect.Slice {
		//fields := iterate(d.MapIndex(k).Interface())
		//d := reflect.ValueOf(data)
		//tmpData := make([]interface{}, d.Len())
		//returnSlice := make([]interface{}, d.Len())
		//for i := 0; i < d.Len(); i++ {
		//	tmpData[i] = d.Index(i).Interface()
		//}
		//
		//for i, v := range tmpData {
		//	returnSlice[i] = iterate(v)
		//}
		//
		//return returnSlice
		d := reflect.ValueOf(data)
		for i := 0; i <d.Len(); i++ {
		dTemp, err := d.Index(i).Interface().(map[string]interface{})
		if !err {
			continue
		}

			for k, v := range dTemp {

					typeOfValue := reflect.TypeOf(v).Kind()

					if typeOfValue == reflect.Map  || typeOfValue == reflect.Slice  {
						fields := iterate(v)
						if isUnique(tmp, k) {
							tmp = append(tmp, schema{k, schemaType(typeOfValue.String()), "NULLABLE", &fields})
						}
					} else {
						if isUnique(tmp, k) {
							tmp = append(tmp, schema{k, schemaType(typeOfValue.String()), "NULLABLE", nil})
						}
					}


			}

		}
	} else if reflect.ValueOf(data).Kind() == reflect.Map {
		d := reflect.ValueOf(data)
		//tmpData := make(map[string]interface{})
		//iterator := d.MapRange()
		//for iterator.Next() {
		//	value := iterator.Value()
		//	tmp= append(tmp,schema{iterator.Key().String(),reflect.ValueOf(value).Kind().String(),"NULLABLE",nil})
		//}



		for _, k := range d.MapKeys() {

			typeOfValue := reflect.TypeOf(d.MapIndex(k).Interface()).Kind()

			if typeOfValue == reflect.Map || typeOfValue == reflect.Slice {
				fields := iterate(d.MapIndex(k).Interface())
				if isUnique(tmp, k.String()) {
					tmp = append(tmp, schema{k.String(), schemaType(typeOfValue.String()), "NULLABLE", &fields})
				}
			} else {
				if isUnique(tmp, k.String()) {
					tmp = append(tmp, schema{k.String(), schemaType(typeOfValue.String()), "NULLABLE", nil})
				}
			}

		}

	}
	return tmp
}
func isUnique(tmp []schema, key string)bool{
	for _, v := range tmp{
		if v.Name == key && v.Type != "RECORD" {
			return false
		}
	}
	return true
}

func schemaType(str string)string{
	switch str {
	case "string":
		return "STRING"
	case "slice":
		return "RECORD"
	case "float64":
		return "FLOAT"
	case "float32":
		return "FLOAT"
	case "int":
		return "INTEGER"
	case "map":
		return "RECORD"
	default:
		return str
	}
}
