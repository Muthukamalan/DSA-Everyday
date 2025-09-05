package main

func getMaxMinInArray(list []int) (int, int) {
	max := list[0]
	min := list[0]
	for _, v := range list {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func checkKInArray(list []int, k int) bool {
	for _, v := range list {
		if v == k {
			return true
		}
	}
	return false
}

func CountElementInArray(list []int, k int) int {
	count := 0
	for _, v := range list {
		if v == k {
			count++
		}
	}
	return count
}

func EvenOddDifference(n []int) int {
	even := 0
	odd := 0
	for i := 0; i < len(n); i++ {
		if n[i]%2 == 0 {
			even += n[i]
		} else {
			odd += n[i]
		}
	}
	return even - odd
}

func FrequencyOfArray(list []int) map[int]int {
	m := make(map[int]int)
	for _, v := range list {
		m[v] += 1
	}
	return m
}

func ConsecutiveSum(lst []int, k int) bool {
	for i := 0; i < len(lst)-1; i++ {
		if lst[i]+lst[i+1] == k {
			return true
		}
	}
	return false
}

func CheckPairIncOrder(lst []int) bool {
	if len(lst) < 2 {
		return true
	} else {
		for i := 0; i < len(lst)-1; i++ {
			if lst[i] > lst[i+1] {
				return false
			}
		}
		return true
	}
}
