package towsum

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {

		idx, ok := m[target-v]
		// fmt.Println(i, v, ok, m[target-v])

		if ok {
			return []int{i, idx}
		} else {
			m[v] = i
		}

		fmt.Println(m)

	}
	return []int{0, 0}
}

func cal(numarray []int, traget int) []int {
	indexarr := []int{}

	for i := 0; i < len(numarray); i++ {
		for x := 0; x < len(numarray); x++ {
			if i != x {
				sum := numarray[i] + numarray[x]
				// fmt.Println(numarray[i], numarray[x])
				if sum == traget {
					indexarr = append(indexarr, i)
				}
			}
		}
	}

	return indexarr
}

// func main() {
// 	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
// 	fmt.Println(twoSum([]int{3, 2, 4}, 6))
// 	fmt.Println(twoSum([]int{3, 3}, 6))

// }
