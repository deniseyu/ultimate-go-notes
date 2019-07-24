package main

import "fmt"

type user struct {
	email string
	name  string
}

func (u user) notify() {
	fmt.Println("hello", u.name)
}

type list []int

func (l list) len() int {
	return len(l)
}

func (l *list) append(val int) {
	*l = append(*l, val)
}

type admin struct {
	person user // not embedding
	level  string
}

type admin2 struct {
	user  // embedding!
	level string
}

func main() {
	var lst list
	lst.append(1)
	fmt.Printf("value receiver list: %v", lst)

	plst := new(list) // returns an initialized pointer
	plst.append(2)
	fmt.Printf("pointer receiver list: %v", plst)

	// embedding

	ad := admin2{
		user:  user{name: "bob", email: "bob@hi.com"},
		level: "super",
	}

	// the following values return the same
	fmt.Println("\n", ad.name) // inner type promotion
	fmt.Println("\n", ad.user.name)
}
