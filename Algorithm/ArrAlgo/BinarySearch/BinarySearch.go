package BinarySearch

/*
	func main() {
		arr := []int{1, 2, 3, 4, 5, 6, 7}
		fmt.Println(BinarySearch(arr, 6))
	}
*/
func BinarySearch(arr []int, target int) int {
	//二分查找法
	var left int = 0
	var right int = len(arr) - 1
	for left <= right {
		module := (left + right) / 2
		if arr[module] > target { //在左边
			right = module - 1
		} else if arr[module] < target { //在右边
			left = module + 1
		} else {
			return module
		}
	}
	//没有找到返回索引-1
	return -1
}
