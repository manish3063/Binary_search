package main

import "fmt"

func main() {
	array := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	data := 9
	var mid int

	l := 0
	r := len(array) - 1
	flag := 0
	for l < r {

		mid = (l + r) / 2
		//fmt.Println(mid)
		if data == array[mid] {
			fmt.Println(array[mid])
			flag = 1
			break

		} else if data < array[mid] {
			r = mid - 1

		} else {
			l = mid + 1

		}

	}
	if flag == 0 {
		fmt.Println("not found")
	}

}
