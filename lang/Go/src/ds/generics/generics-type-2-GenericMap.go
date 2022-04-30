package main

import (
	"fmt"
)

type GenericStringMap[T any] map[string]T

func (g *GenericStringMap[T]) Set(key string, value T) {
	(*g)[key] = value
}

func (g *GenericStringMap[T]) Get(key string) T {
	return (*g)[key]
}

func (g *GenericStringMap[T]) Print() {
	for k, v := range *g {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func Print[T any](g GenericStringMap[T]) {
	for k, v := range g {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func main() {
	g := GenericStringMap[int]{"a": 1}
	g["b"] = 2
	g.Set("c", 3)
	fmt.Println(g.Get("c"))
	fmt.Println()

	g.Print()
	fmt.Println()
	Print(g)
	fmt.Println()

	g2 := GenericStringMap[string]{"A": "foo"}
	g2.Set("B", "bar")
	g2.Print()
	g2.Set("C", "fur")
	fmt.Println("C:", g2.Get("C"))
}
