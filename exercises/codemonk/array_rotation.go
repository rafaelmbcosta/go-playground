package main

import (
	"fmt"
)

func main() {
	var test_cases int
	fmt.Scanf("%d", &test_cases)

	for cases := 0; cases < test_cases; cases++ {
		var n, k int //, n, k int
		fmt.Scanf("%d %d", &n, &k)
		array := make([]int, n)

		for i := 0; i < n; i++ {
			var input int
			fmt.Scanf("%d", &input)

			var rotations = k % n
			if (i + rotations) > n-1 {
				array[(i+rotations)-n] = input
			} else {
				array[i+rotations] = input
			}
		}

		for _, value := range array {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
