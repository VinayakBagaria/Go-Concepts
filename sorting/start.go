package sorting

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99)
	}
	return slice
}

func DoWork() {
	slice := generateSlice(10)
	fmt.Println(slice)
	fmt.Println("-----------")

	fmt.Println("Insertion Sort")
	fmt.Println(InsertionSort(slice))
	fmt.Println("-----------")

	fmt.Println("Selection Sort")
	fmt.Println(SelectionSort(slice))
	fmt.Println("-----------")

	fmt.Println("Merge Sort")
	fmt.Println(MergeSort(slice))
	fmt.Println("-----------")

}
