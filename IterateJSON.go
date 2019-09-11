package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type schema struct {
	Name   string    `json:"name"`
	Type   string    `json:"type"`
	Mode   string    `json:"mode"`
	Fields *[]schema `json:"fields,omitempty"`
}

func iterate(data interface{}) []schema {

	tmp := make([]schema, 0)

	if reflect.ValueOf(data).Kind() == reflect.Slice {

		d := reflect.ValueOf(data)
		for i := 0; i < d.Len(); i++ {
			dTemp, err := d.Index(i).Interface().(map[string]interface{})
			if !err {
				continue
			}

			for k, v := range dTemp {

				typeOfValue := reflect.TypeOf(v).Kind()

				if typeOfValue == reflect.Map || typeOfValue == reflect.Slice {
					fields := iterate(v)

					tmp = append(tmp, schema{k, schemaType(v), isArray(typeOfValue.String()), fieldz(&fields)})

				} else {

					tmp = append(tmp, schema{k, schemaType(v), isArray(typeOfValue.String()), nil})

				}

			}

		}
	} else if reflect.ValueOf(data).Kind() == reflect.Map {
		d := reflect.ValueOf(data)

		for _, k := range d.MapKeys() {

			typeOfValue := reflect.TypeOf(d.MapIndex(k).Interface()).Kind()

			if typeOfValue == reflect.Map || typeOfValue == reflect.Slice {
				fields := iterate(d.MapIndex(k).Interface())

				tmp = append(tmp, schema{k.String(), schemaType(d.MapIndex(k).Interface()), isArray(typeOfValue.String()), &fields})

			} else {

				tmp = append(tmp, schema{k.String(), schemaType(d.MapIndex(k).Interface()), isArray(typeOfValue.String()), nil})

			}

		}

	}

	return tmp
}

func fieldz(fields *[]schema) *[]schema {
	if len(*fields) > 0 {
		return fields
	}
	return nil
}
func isArray(str string) string {
	switch str {
	case "slice":
		return "REPEATED"
	case "map":
		return "REPEATED"
	default:
		return "NULLABLE"
	}
}

func schemaType(value interface{}) string {
	myType := reflect.TypeOf(value)
	if n, ok := value.(json.Number); ok {
		// myVar was a number, let's see if its float64 or int64
		// Check for int64 first because floats can be parsed as ints but not the other way around
		if v, err := n.Int64(); err != nil {
			// The number was an integer, v has type of int64
			fmt.Println(v)
		}
		if v, err := n.Float64(); err != nil {
			// The number was a float, v has type of float64
			fmt.Println(v)
		}
	} else {
		// myVar wasn't a number at all
	}
	switch myType.Kind() {

	case reflect.String:
		return "STRING"
	case reflect.Slice:
		return "RECORD"
	case reflect.Float64:
		return "FLOAT"
	case reflect.Float32:
		return "FLOAT"
	case reflect.Int:
		return "INTEGER"
	case reflect.Int64:
		return "INTEGER"
	case reflect.Map:
		return "RECORD"
	default:
		return myType.Kind().String()
	}

}