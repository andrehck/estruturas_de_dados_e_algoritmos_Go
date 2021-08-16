package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var lenght = len(scores)

	sort(scores, lenght)

	for i := 0; i < lenght; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func sort(arrays []int, lenght int) {
	var minIndex int

	for i := 0; i < lenght-1; i++ {
		minIndex = i
		//O indice do minimo selecionado

		var minValue = arrays[minIndex]
		for j := i; j < lenght-1; j++ {
			if minValue > arrays[j+1] {
				minValue = arrays[j+1]
				minIndex = j + 1
			}
		}
		//O minimo de arrays[i] é trocado pelos arrays [minIndex]
		if i != minIndex {
			var temp = arrays[i]
			arrays[i] = arrays[minIndex]
			arrays[minIndex] = temp
		}
	}
}
