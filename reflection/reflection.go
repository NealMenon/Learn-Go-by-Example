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
	field := val.Field(0)
	fn(field.String())
}

func main() {
	fmt.Println("Build")
}
