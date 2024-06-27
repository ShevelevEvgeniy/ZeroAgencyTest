package repository_query

import (
	"reflect"
	"strings"
)

func ColumnsToUpdate(obj interface{}) []string {
	var columns []string
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		reformTag := typ.Field(i).Tag.Get("reform")

		columnName := strings.Split(reformTag, ",")[0]

		if columnName == "id" {
			continue
		}

		if columnName != "" {
			columns = append(columns, columnName)
		}
	}

	return columns
}
