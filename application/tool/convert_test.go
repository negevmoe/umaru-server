package tool

import (
	"fmt"
	"testing"
)

func TestArrayDeDuplicate(t *testing.T) {
	a := []string{"a", "a", "a", "b", "c", "c", "d", "", ""}
	b := []int64{1, 0, 2323, 0, 1, 123, 4444, -2, -2, -1, -6}
	ArrayDeDuplicate(&a)
	ArrayDeDuplicate(&b)
	for _, item := range a {
		fmt.Println("item:", item)
	}
	fmt.Println("*****************")
	for _, item := range b {
		fmt.Println("item:", item)
	}
}

func TestArray2Set(t *testing.T) {
	a := []string{"a", "a", "a", "b", "c", "c", "d", "", ""}
	b := []int64{1, 0, 2323, 0, 1, 123, 4444, -2, -2, -1, -6}
	aset := Array2Set(a)
	bset := Array2Set(b)
	fmt.Println(aset)
	fmt.Println("******************")
	fmt.Println(bset)
}
