package main

import "fmt"

func main() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)

	type cat struct {
		mew string
		age int32
	}

	// anonymous structs, or literal types
	e1 := struct {
		flag    bool
		counter int64
	}{
		flag:    true,
		counter: 10,
	}

	fmt.Printf("e1: %v\n", e1)

	// passing by value
	f := 10
	fmt.Printf("f value: %v \t address %v\n", f, &f)

	num := 0
	increment(&num)
	fmt.Println(num)
}

func increment(i *int) {
	*i++
}
