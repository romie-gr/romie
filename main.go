/*
Copyright 2020 Romie Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install github.com/client9/misspell/cmd/misspell
//go:generate go install golang.org/x/tools/cmd/goimports

package main

import (
	"fmt"

	"github.com/romie-gr/romie/internal/debug"
)

func main() {
	println("Hello romie")

	myArray2 := []int{2, 4, 5, 6, 7, 8}
	myArrayf := []string{"kolos"}

	//debug.PrintArray(myArray2)
	// debug.ArrayAllTypes(debug.Print, myArrayf)
	// debug.ArrayAllTypes(debug.Print, myArray2)
	debug.PrintArrayReflect(myArray2)
	debug.PrintArrayReflect(myArrayf)

	fmt.Println(debug.GetTypeArray(myArray2))
	fmt.Println(debug.GetTypeArray(myArrayf))
}

func Summary(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
