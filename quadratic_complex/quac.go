package main

import "fmt"

// Se representa como O(N2) y el procesamiento
// es proporcional al numero de inputs al cuadrado

func main() {
	var k, l int
	for k = 0; k <= 10; k++ {
		fmt.Println("Multiplication Table", k)
		for l = 1; l <= 10; l++ {
			var x int = l * k
			fmt.Println(x)
		}
	}
}
