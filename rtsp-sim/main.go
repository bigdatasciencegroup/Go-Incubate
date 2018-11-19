package main

import "fmt"

func main() {

	g := []int{1,2,3}
	hi(&g)
	fmt.Print(g)

}

func hi(k *[]int){
	*k = append(*k, 0)
}
