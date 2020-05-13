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
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
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

func main() {
	fmt.Println("Build")
}
