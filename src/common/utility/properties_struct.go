package utility

import "reflect"

func GetNamedStruct(data interface{}) []string {
	var value []string
	val := reflect.ValueOf(data)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Type().Field(i).Tag.Get("db") == "" {
			continue
		}
		value = append(value, val.Type().Field(i).Tag.Get("db"))
	}
	return value
}
