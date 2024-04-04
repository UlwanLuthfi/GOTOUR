package main

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func TestWalk(t *testing.T) {
	max := 10
	want := make([]int, 0)
	for i := 1; i <= max; i++ {
		for j := 1; j <= i; j++ {
			want = append(want, j*i)
		}
	}

	tree := tree.New(1)
	ch := make(chan int)
	go Walk(tree, ch)
	for i := 0; i < len(want); i++ {
		v := <-ch
		if want[i] != v {
			t.Errorf("got %d, want %d", v, want[i])
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1, v2 := <-ch1, <-ch2
		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
