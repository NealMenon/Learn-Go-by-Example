package main

import (
	"fmt"
	"reflect"
)

/*
 golang challenge: write a function
 walk(x interface{}, fn func(string))
 which takes a struct x and calls fn for
 all strings fields found inside.
 difficulty level: recursively.
*/
func walk(x interface{}, fn func(string)) {

	val := getValue(x) // catches if thrown a pointer

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}

	}

}

// getValue catches if thrown a pointer
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func main() {
	fmt.Println("Build")
}
