package main

import "fmt"

func main() {
	test := []string{"()())(", "()", "{}[[][]]"}
	for _, t := range test {
		fmt.Println(t, ":", isValid(t))
	}

}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	m := map[string]string{")": "(", "]": "[", "}": "{"}
	stack := []string{}
	for _, w := range s {
		if _, ok := m[string(w)]; !ok {
			stack = append(stack, string(w))
		} else if len(stack) == 0 {
			return false
		} else {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if m[string(w)] != popped {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
