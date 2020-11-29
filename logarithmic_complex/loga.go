package main

import (
	"fmt"
	"time"
)

type Tree struct {
	LeftNode  *Tree
	Value     int
	RigthNode *Tree
}

// Por lo general un logaritmo siempre sera proporcional al lagoritmo del numero de inputs
// La complejidad de la insercion de entrada es de tipo O(log N)
func (tree *Tree) insert(m int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RigthNode == nil {
				tree.RigthNode = &Tree{nil, m, nil}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RigthNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil, m, nil}
	}
}

func print(tree *Tree) {
	if tree != nil {
		fmt.Println("Value", tree.Value)
		fmt.Printf("Tree Node Left")
		print(tree.LeftNode)
		fmt.Printf("Tree Node Rigth")
		print(tree.RigthNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

func main() {
	start := time.Now()
	var tree *Tree = &Tree{nil, 1, nil}
	print(tree)
	tree.insert(3)
	print(tree)
	tree.insert(5)
	print(tree)
	tree.LeftNode.insert(7)
	print(tree)
	tree.RigthNode.insert(8)
	//tree.RigthNode.insert(15)
	print(tree)
	//tree.RigthNode.LeftNode.insert(57) si le agregamos se aumenta el tiempo a 20 por que 5^2
	//print(tree)
	t := time.Now()
	time := t.Sub(start)
	fmt.Println(time)
}
