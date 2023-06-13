package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Array e Slices")

	var array1 [5]int
	array1[0] = 15
	fmt.Println(array1)

	array2 := [...]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)

	slice := []int{10, 2, 6, 9, 8, 10}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, 56)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Alterado"
	fmt.Println(slice2)

	fmt.Println("-----------")

	// Arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) // capacidade

}
