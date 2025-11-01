package RemoveElement

func RemoveElement[T comparable](slice []T, target T) []T {
	//移除元素	1.暴力解法O(n^2)
	k := 0
	for k < len(slice) {
		if slice[k] == target {
			for i := k + 1; i < len(slice); i++ {
				slice[i-1] = slice[i]
			}
			slice = slice[:len(slice)-1]
		} else {
			k++
		}
	}
	return slice
}
