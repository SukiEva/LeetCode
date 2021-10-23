package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		if v == "+" {
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		} else if v == "-" {
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])
		} else if v == "*" {
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		} else if v == "/" {
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
		} else {
			t, _ := strconv.Atoi(v)
			stack = append(stack, t)
		}
	}
	return stack[0]
}
