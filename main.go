package main

import (
	"fmt"
	"github.com/steveoc64/memdebug"
	"os"
	"time"
)

type fourSet [4]int
type fourMap map[fourSet]struct{}

func permutations(arr []int) fourMap {
	var helper func([]int, int)
	//res := []fourSet{}
	significantLength := len(arr) / 2
	if significantLength > 4 {
		significantLength = 4
	}
	t1 := time.Now()
	memdebug.Print(t1, "output array is uniques of size", significantLength)
	tmp := fourSet{}

	res2 := make(fourMap)

	helper = func(arr []int, n int) {
		if n == 1 {
			//copy(tmp, arr)
			for i := 0; i < significantLength; i++ {
				tmp[i] = arr[i]
			}
			res2[tmp] = struct{}{}

			/*
				alreadyThere := false
				for i := 0; i < len(res); i++ {
					if res[i] == tmp {
						alreadyThere = true
						break
					}
				}
				if !alreadyThere {
					res = append(res, fourSet{
							tmp[0],tmp[1],tmp[2],tmp[3],
						})


				} else {
					//memdebug.Print(t1, "IS ALREADY THERE", tmp)
				}
			*/
		} else {
			flipflop := true
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

	// sort the results
	/*
		sort.Slice(res, func(i, j int) bool {
			for k := 0; k < significantLength; k++ {
				if res[i][k] < res[j][k] {
					return true
				}
				if res[i][k] > res[j][k] {
					return false
				}
				// else they are equal - try the next digit
			}
			return false
		})
	*/
	return res2
}

func main() {
	memdebug.Profile()
	defer memdebug.WriteProfile()

	debug := os.Getenv("QUIET") == ""
	t1 := time.Now()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	memdebug.Print(t1, "start", arr)
	perm := permutations(arr)
	memdebug.Print(t1, "reduced permutations to", len(perm))
	t1 = time.Now()
	if debug {
		i := 0
		for k := range perm {
			if i > 0 {
				if i%16 == 0 {
					fmt.Println("")
				} else {
					fmt.Print(", ")
				}
			}
			i++
			fmt.Print(k)
		}
		fmt.Println("")
	}
	memdebug.Print(t1, "done results", len(perm))
}
