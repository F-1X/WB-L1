package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.
func main() {

	var array []interface{}
	array = append(array, 1, "string", true, make(chan int))


	fmt.Println("defineTypeSwitching:\n-----------")
	defineTypeSwitching(array)
	fmt.Println("defineReflection:\n-----------")
	defineReflection(array)

}

func defineTypeSwitching(array []interface{}) {
	for i := 0; i < len(array); i++ {
		// type switching
		switch array[i].(type) {
		case int:
			fmt.Println("int:", array[i].(int))
		case string:
			fmt.Println("string:", array[i].(string))
		case bool:
			fmt.Println("bool:", array[i].(bool))
		case chan int:
			fmt.Println("chan int:", array[i].(chan int))
		default:
			fmt.Println("unknown")
		}

	}
}

func defineReflection(array []interface{}) {
	for i := 0; i < len(array); i++ {
		// рефлексия
		switch reflect.TypeOf(array[i]).Kind() {
		case reflect.Int:
			fmt.Println("int:",reflect.ValueOf(array[i]))
		case reflect.String:
			fmt.Println("string:",reflect.ValueOf(array[i]))
		case reflect.Bool:
			fmt.Println("bool:",reflect.ValueOf(array[i]))
		case reflect.Chan:
			fmt.Println("chan int:",reflect.ValueOf(array[i]))
		default:
			fmt.Println("unknown")
		}
	}
}
