package main

import (
	"fmt"
	"save/study/Algorithm/ArrAlgo/RemoveElement"
)

func main() {
	var arr = []int{2, 1, 2, 2, 3, 2, 4, 5, 7, 2}
	newArr := RemoveElement.RemoveElement(arr, 2)
	fmt.Println(newArr)
}
