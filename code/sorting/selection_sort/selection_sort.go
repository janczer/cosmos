package main

import "fmt"

func main() {
	a := []int{87, 22, 31, 11, 3, 98, 12, 11}
	fmt.Println(a)
	selectionSort(a)
	fmt.Println(a)
}

func selectionSort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		aMin := a[i]
		iMin := i
		for j := i + 1; j < n; j++ {
			if a[j] < aMin {
				aMin = a[j]
				iMin = j
			}
		}
		a[i], a[iMin] = aMin, a[i]
	}
}
