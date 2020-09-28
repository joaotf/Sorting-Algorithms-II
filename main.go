package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func perguntaproDoidao() int {
	var tamanho int
	fmt.Print("\n\nTamanho da ARRAY : ")
	_, err := fmt.Scan(&tamanho)
	if err == nil {
		fmt.Println()
	}
	return tamanho
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	var tamanho int
	var array []int
	var menu int

	var list []int
	var countx int
	var tempo float64

	fmt.Print("Menu\n1) Quick Sort\n2) Merge Sort\n3) Heap Sort\nOpção:")
	_, err := fmt.Scan(&menu)
	if err == nil {
	}
	for menu != 0 {
		switch menu {
		case 1:
			{
				tamanho = perguntaproDoidao()
				array = generateSlice(tamanho)
				fmt.Println("Array não ORDENADO : ", array)
				list, countx, tempo = quicksort(array)

				fmt.Println("\n\nQuick Sort (Finished) : ", list, "\n\nContador : ", countx)
				fmt.Printf("Tempo : %.20f", tempo)
				os.Exit(3)
			}

		case 2:
			{
				tamanho = perguntaproDoidao()
				array = generateSlice(tamanho)
				fmt.Println("Array não ORDENADO : ", array)
				list = mergeSort(array)
				fmt.Println("\n\nMerge Sort (Finished) : ", list, "\n\nContador : ", count)
				fmt.Printf("\nTempo : %.20f", -tempofinal)
				os.Exit(3)
			}

		case 3:
			{
				tamanho = perguntaproDoidao()
				array = generateSlice(tamanho)
				fmt.Println("Array não ORDENADO : ", array)
				list = heapSort(array)
				fmt.Println("\n\nBubble Sort (Finished) : ", list, "\n\nContador : ", contador)
				fmt.Printf("\nTempo : %.20f", -tempoheap)
				os.Exit(3)
			}

		default:
			fmt.Println("Opção Inválida!")
			break

		}

	}
}
