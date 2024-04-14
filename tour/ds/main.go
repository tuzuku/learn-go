package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x int
	y int
}

func main() {
	defer fmt.Println("end....")
	fmt.Println("start....")

	v := Vertex{x: 1}
	fmt.Println(v)

	p := &v
	p.y = 200
	fmt.Println(v)

	for i := 0; i < 10; i++ {
		fmt.Print("i=", i, " ")
	}
	for v.x < 10 {
		v.x++
		fmt.Print("v.x=", v.x, " ")
	}

	fmt.Println()
	arr := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(arr)

	var s []int = arr[1:4]
	fmt.Println(s)

	r := []struct {
		val int
		b   bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, false},
		{11, true},
		{13, false},
	}
	fmt.Println(r)
	fmt.Println(len(r), cap(r))

	arr2 := make([]int, 2, 5)

	fmt.Println(arr2)
	arr2[0] = 100
	fmt.Println(arr2[0])
	arr2 = append(arr2, 200)
	fmt.Println(arr2)
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	for i := range arr2 {
		fmt.Println(i)
	}

	for _, v := range arr2 {
		fmt.Println(v)

	}

	m := make(map[string]Vertex)
	m["zozo"] = Vertex{3, 4}
	fmt.Println(m)
	fmt.Println(m["zozo"])
	fmt.Println(m["zozo"].x)

	m["yoyo"] = Vertex{5, 6}
	fmt.Println(m)

	fmt.Println(m["yoyo"].x)
	fmt.Println(m)
	delete(m, "yoyo")
	fmt.Println(m)

	add := func(a, b float64) float64 {
		return a + b
	}
	add(2, 4)
	fmt.Println(compute(add))

	fmt.Println(compute(math.Pow))

}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
		for j := range p[i] {
			p[i][j] = uint8(i * j)
		}
	}
	return p
}
