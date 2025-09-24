package main

func MaxModInArray(list []int) int {
	mod := 0
	length := len(list)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {

			cmod := list[i] % list[j]
			if cmod > mod {
				mod = cmod
			}
		}
	}
	return mod
}
