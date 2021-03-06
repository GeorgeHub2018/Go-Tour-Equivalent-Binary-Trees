package main

import "golang.org/x/tour/tree"
import "fmt"

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	go Walk(t1, ch1)
	ch2 := make(chan int, 10)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		if x != y {
			return false
		}
	}

	return true
}

func main() {
	// test walk
	ch := make(chan int, 10)
	t := tree.New(1)
	go Walk(t, ch)
	// print walk
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	// test same
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
