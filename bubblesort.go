package main

import "fmt"

func main() {
	var score = []int{90, 70, 50, 80, 60, 85}
	var lenght = len(score)

	sort(score, lenght)

	for i := 0; i < lenght; i++ {
		fmt.Printf("%d ", score[i])
	}
}

func sort(arrays []int, lenght int) {
	for i := 0; i < lenght-1; i++ {
		var isSwap = false
		for j := 0; j < lenght-i-1; j++ {
			if arrays[j] > arrays[j+1] {
				var flag = arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = flag
				isSwap = true
			}
		}
		if isSwap == false {
			break
		}
	}

}
