package slicer

import "fmt"

type User struct {
	Name string
	Age  int16
	Mail string
}

func SliceTest() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := [6]bool{true, false, true, true, false, true}
	fmt.Println(t)

	b := make([]int8, 5)
	bb := make([]bool, 1)
	ints := append(b, 1, 2, 3, 4, 5)

	for _, value := range ints {
		fmt.Println(value)
	}

	bb = append(bb, true, false)

	fmt.Println(b, ints)

	type User struct {
		Name string
	}

	pointer := &User{Name: "Martin"}
	pointer.Name = "Harald"
	fmt.Println(pointer)

	asPointer := getUserAsPointer("Harald")
	fmt.Printf("Pointer %p", asPointer)

	var gloriousSentencesOfMartin = map[string]string{}
	gloriousSentencesOfMartin["martin"] = "Hello World"
	gloriousSentencesOfMartin["martin2"] = "how much is the fish"
	fmt.Println(gloriousSentencesOfMartin)
}

func getUserAsPointer(name string) *User {
	return &User{Name: name}
}
