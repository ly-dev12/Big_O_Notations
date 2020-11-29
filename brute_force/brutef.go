package main

import "fmt"

// Estos algoritmos suelene ser nromalmente lentos y no tan eficientes
// Se pueden utilizar para la busqueda, la coincidencia de cadenas y la multiplicacion de amtrices
func findElement(arr [10]int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}

	return false
}
func main() {
	var arr = [10]int{1, 7, 8, 3, 5, 9, 4, 2, 6}
	var check bool = findElement(arr, 10)
	fmt.Println(check)
	var check2 bool = findElement(arr, 9)
	fmt.Println(check2)
}
