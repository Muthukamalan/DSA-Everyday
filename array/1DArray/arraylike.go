package main

func InsertAtEnd(list *[]int, val int) []int {
	copyList := make([]int, len(*list)+1)
	copy(copyList, *list)
	copyList[len(copyList)-1] = val
	return copyList
}

func InsertAtBegin(list *[]int, val int) []int {
	copyList := make([]int, len(*list)+1)
	copyList[0] = val
	copy(copyList[1:], *list)
	return copyList
}

func RightShiftArray(a []int) []int {
	n := len(a)
	if n == 0 {
		return a
	} else {
		last := a[n-1]
		for i := n - 1; i > 0; i-- {
			a[i] = a[i-1]
		}
		a[0] = last
		return a
	}
}
