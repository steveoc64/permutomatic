package main

import (
	"fmt"
	"github.com/steveoc64/memdebug"
	"os"
	"time"
)

func sliceSame(a,b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}
	flipflop := true
	significantLength := len(arr) / 2
	t1 := time.Now()
	memdebug.Print(t1, "sigLen", significantLength)

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, significantLength)
			//copy(tmp, arr)
			for i := 0; i < significantLength; i++ {
				tmp[i] = arr[i]
			}
			//memdebug.Print(t1, "adding this array here", tmp, arr)
			// we now have one to add - check to see if its already there
			alreadyThere := false
			for i := 0; i < len(res); i++ {
				if sliceSame(res[i], tmp) {
					alreadyThere = true
				}
			}
			if !alreadyThere {
				res = append(res, tmp)
			} else {
				//memdebug.Print(t1, "IS ALREADY THERE", tmp)
			}
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if flipflop {
					arr[0], arr[n-1] = arr[n-1], arr[0]
					flipflop = false
				} else {
					arr[i], arr[n-1] = arr[n-1], arr[i]
					flipflop = true
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	debug := os.Getenv("DEBUG") != ""
	t1 := time.Now()
	memdebug.Print(t1, "start")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("// will generate [1 2 3 4],[2 1 3 4],[3 2 1 4]")
	perm := permutations(arr)
	memdebug.Print(t1, "done permutations", len(perm))
	t1 = time.Now()
	if debug {
		for i := 0; i < int(len(perm)); i++ {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(perm[i])
		}
		fmt.Println("")
	}
	memdebug.Print(t1, "done print results", len(perm))
	fmt.Println("// Would like to generate only pairs of two and most lovely unique")
	t1 = time.Now()
	if debug {
		for i := 0; i < int(len(perm)); i++ {
			if i > 0 {
				if i%24 == 0 {
					fmt.Println()
				} else {
					fmt.Print(", ")
				}
			}
			for j := 0; j < len(arr)/2; j++ {
				fmt.Print(perm[i][j])
			}
		}
	}
	fmt.Println("")
	memdebug.Print(t1, "all done")

	// simplify results


}
