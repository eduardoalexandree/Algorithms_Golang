package main

import "fmt"

func searchMin(arr []int) int {

	minIndex := 0
	minValue := arr[0]
	for i, _ := range arr[1:] {
		if arr[i] < minValue {
			minValue = arr[i]
			minIndex = i
		}
	}
	return minIndex
}

func selectionSort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)

	for i, _ := range arr {
		min := searchMin(arr)
		newArr[i] = arr[min]
		arr = append(arr[:min], arr[min+1:]...)
	}
	return newArr
}

func main() {
	arr := []int{5, 4, 3, 2, 1, 0}
	fmt.Println(selectionSort(arr))
}
