package Recursion

func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		return -1
	}

	mid := (l + r) / 2
	if target == array[mid] {
		return mid
	} else if target < array[mid] {
		return BinarySearch(array, target, l, mid-1)
	} else {
		return BinarySearch(array, target, mid+1, r)
	}
}
