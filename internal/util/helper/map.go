package helper

import "reflect"

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		// Check if the field is exported
		if field.PkgPath == "" {
			value := val.Field(i).Interface()

			if valueKind := val.Field(i).Kind(); valueKind == reflect.Struct {
				// Recursively convert nested structs
				result[field.Name] = StructToMap(value)
			} else {
				result[field.Name] = value
			}
		}
	}
	return result
}
