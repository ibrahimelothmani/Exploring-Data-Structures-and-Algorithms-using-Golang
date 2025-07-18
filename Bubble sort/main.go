package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main(){
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(bubbleSort(arr))
}