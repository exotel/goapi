package helpers

import (
	"reflect"
	"time"
)

//StructToMap converts one or more struct to a map of string to interface
//Its uses the tag 'tag' to check the name to be used ,if the second struct had the same
//tag name the second one will be overriding the first
//if the value is nil or no tag is defined no key is added
func StructToMap(tag string, structs ...interface{}) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for _, in := range structs {
		v := reflect.ValueOf(in)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		// we only accept structs
		if v.Kind() != reflect.Struct {
			continue
		}
		tp := v.Type()
		for i := 0; i < tp.NumField(); i++ {
			field := tp.Field(i)
			if tagv := field.Tag.Get(tag); tagv != "" && !IsEmptyValue(v.Field(i)) {
				m[tagv] = v.Field(i).Interface()
			}
		}
	}
	return
}

//IsEmptyValue checks if the reflect.Value given is empty or not
func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		switch v.Type().Name() {
		case "Time":
			return v.Interface().(time.Time) == (time.Time{})
		}

	}
	return false
}
