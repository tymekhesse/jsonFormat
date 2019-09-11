package main

import (
	"encoding/json"
	"reflect"
	"strings"
)

type schema struct {
	Name   string    `json:"name"`
	Type   string    `json:"type"`
	Mode   string    `json:"mode"`
	Fields *[]schema `json:"fields,omitempty"`
}

func iterate(data interface{}) []schema {
	tmp := make([]schema, 0)
	d := reflect.ValueOf(data)

	if d.Kind() == reflect.Slice {
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
	} else if d.Kind() == reflect.Map {
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
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return "STRING"
	case reflect.Slice, reflect.Map:
		return "RECORD"
	case reflect.Float64:
		x, _ := json.Marshal(value)
		if strings.Contains(string(x), ".") {
			return "FLOAT"
		}
		return "INTEGER"
	default:
		return reflect.TypeOf(value).Kind().String()
	}
}
