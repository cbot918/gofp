package main

import "fmt"

type chain struct {
	V int
}

func from(v int) chain {
	return chain{V: v}
}

func (c chain) then(fn func(int) int) chain {
	c.V = fn(c.V)
	return c
}

func (c chain) ret() int {
	return c.V
}

func main() {

	answer := from(3).
		then(func(x int) int { return x + 1 }).
		then(func(x int) int { return x * 5 }).
		ret()

	fmt.Println(answer)
}
