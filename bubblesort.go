package main

import (
	"fmt"
)

func main() {
	var scores = []int{1, 6, 3, 8, 4, 9, 0, 3}
	var lenght = len(scores)
	fmt.Println(sort(scores, lenght))

}

func sort(array []int, lenght int) []int {
	for i := 0; i < lenght-1; i++ { //adiciona 1 até i ser igual a 7
		var isSwap = false

		for j := 0; j < lenght-i-1; j++ { // i=0;j=0;tamanhodoarray-i-1
			if array[j] > array[j+1] { //comparação entre o numero do array e seu proximo
				var flag = array[j] // se for maior a variavel flag é igual a posiçao do array com index j
				fmt.Println(flag)
				array[j] = array[j+1] // a posicao j é alterada pelo j+1(proximo numero)
				fmt.Println(flag)
				array[j+1] = flag // a posicão que seria menor é alterada pela flag
				fmt.Println(flag)
				isSwap = true // continua realizando a troca
			}
		}
		if isSwap == false {
			break
		}
	}
	return array
}
