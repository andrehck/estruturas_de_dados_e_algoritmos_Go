package insertsort

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
	for i := 0; i < lenght; i++ {
		var insertElement = arrays[i] //Pegue novos elementos não classificados

		var insertPosition = i
		for j := insertPosition - 1; j >= 0; j-- {
			//insertElement é deslocado para a direita

			if insertElement < arrays[j] {
				arrays[j+1] = arrays[j]
				insertPosition--
			}
		}
		arrays[insertPosition] = insertElement
	}
}
