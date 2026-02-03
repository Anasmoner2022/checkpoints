package piscine

func Unmatch(a []int) int {
	a = sort(a)
	for i := 0; i < len(a)-1; i += 2 {
		if a[i] != a[i+1] {
			return a[i]
		}
	}
	if len(a)%2 != 0 {
		return a[len(a)-1]
	}
	return -1
}

func sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	return a
}
