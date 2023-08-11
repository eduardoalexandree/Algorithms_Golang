package main

import "fmt"

func main() {
	list := []int{14, 15, 16, 17, 18, 19, 20, 21}
	value := 18
	if binarySearch(list, value) != nil {
		fmt.Println("Value found")
	} else {
		fmt.Println("Value not found")
	}
}

func binarySearch(l []int, value int) *int {
	low := 0
	high := len(l) - 1
	for low <= high {
		mean := (low + high) / 2
		fmt.Println(l[mean])
		if l[mean] == value {
			return &l[mean]
		}
		if l[mean] > value {
			high = mean - 1
		}
		if l[mean] < value {
			low = mean + 1
		}
	}
	return nil

}
