package Common

func BinarySearch(array []int, target int, l, r int) int {
	ltemp := l
	rtemp := r
	for {
		if ltemp > rtemp {
			return -1
		}

		mid := (ltemp + rtemp) / 2
		midValue := array[mid]
		if midValue == target {
			return mid
		} else if midValue < array[ltemp] {
			rtemp = mid - 1
		} else {
			ltemp = mid + 1
		}
	}
}
