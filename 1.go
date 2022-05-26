package main

import "fmt"

type Human struct {
	name	string
	age		int
}

// Action is a struct which embedds Human
type Action struct {
	Human
}

// Action2 is a struct which has a named Human field, therefore needs promoting
type Action2 struct {
	performer Human
}

func (h Human) introduce() {
	fmt.Printf("My name is %s. I'm %d years old\n", h.name, h.age)
}

// Introduce func on Action2 promotes the method of Human field
func (a Action2) introduce() {
	a.performer.introduce()
}

func main() {
	petr := Human{"Petr", 22}
	petr.introduce()

	a := Action{Human{"Mary", 23}}
	a.introduce()

	a2 := Action2{Human{"Alex", 24}}
	a2.performer.introduce()
	a2.introduce()
}
