package main

func reverseLeftWords(s string, n int) string {
	reserve := func(s []byte) {
		for left, right := 0, len(s)-1; left < right; left++ {
			s[left], s[right] = s[right], s[left]
			right--
		}
	}
	t := []byte(s)
	reserve(t[0:n])
	reserve(t[n:(len(t))])
	reserve(t[0:(len(t))])
	return string(t)
}
