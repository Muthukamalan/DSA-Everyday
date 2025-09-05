package main

func iSArraySorted(l []int) bool {
	length := len(l)
	for i := 0; i < length-1; i++ {
		if !(l[i] <= l[i+1]) {
			return false
		}
	}
	return true

}
