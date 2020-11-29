package main

import (
	"fmt"
)

// Se representa como O(N) y el procesamiento es proporcional al numero de inputs
func main() {
	var m [10]int
	var k int
	for k = 0; k < 10; k++ {
		m[k] = k * 200
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}
