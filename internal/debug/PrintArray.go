package debug

import (
	"fmt"
	"reflect"
)

type Arraytype interface {
}

func ArrayAllTypes(mapper func(Arraytype), list ...Arraytype) {
	for _, value := range list {
		mapper(value)
	}
}

func PrintArray(myArray []int) {
	fmt.Printf("Length: %d\n", len(myArray))
	for i, v := range myArray {
		fmt.Printf("[%d] = [%#v]\n", i, v)
	}
}

func PrintArrayReflect(myArray interface{}) {
	fmt.Printf("real array: %v\n", myArray)
}

func Print(value Arraytype) {
	fmt.Print(value)
}

func GetTypeArray(arr interface{}) reflect.Type {
	return reflect.TypeOf(arr).Elem()
}
