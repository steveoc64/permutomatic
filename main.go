package main

import (
	"fmt"
	"github.com/steveoc64/memdebug"
	"time"
)


func permutations(arr []int)[][]int{
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int){
		if n == 1{
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++{
				helper(arr, n - 1)
				if n % 2 == 1{
					arr[i],arr[n-1] = arr[n-1],arr[i]
				} else {
					arr[0],arr[n-1] = arr[n-1],arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	t1 := time.Now()
	memdebug.Print(t1, "start")
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("// will generate [1 2 3 4],[2 1 3 4],[3 2 1 4]")
	perm := permutations(arr)
	for i := 0; i < int(len(perm)); i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(perm[i])
	}
	fmt.Println("")
	memdebug.Print(t1, "done permute")
	fmt.Println("")
	fmt.Println("// Would like to generate only pairs of two and most lovely unique")
	for i := 0; i < int(len(perm)); i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(perm[i][0], perm[i][1])
	}
	fmt.Println("")
	memdebug.Print(t1, "all done")

}

