package main

func PairCount(l []int, x int) int {
	var length int = len(l)
	counter := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if l[i]+l[j] == x {
				counter += 1
			}
		}
	}
	return counter

}
