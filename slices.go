package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	q[0] = 3
	printSlice(q)
	q = append(q, 4)
	printSlice(q)

	s2 := []int{2, 3, 5, 7, 11, 13}
	s2 = s2[1:4]
	fmt.Println(s2)
	s2 = s2[:2]
	fmt.Println(s2)
	s2 = s2[1:]
	fmt.Println(s2)
	printSlice(s2)

	a2 := make([]int, 5)
	printSlice("a2", a2)
	b2 := make([]int, 0, 5) // len(b2)=0, ca2p(b2)=5
	printSlice("b2", b2)
	c := b2[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
