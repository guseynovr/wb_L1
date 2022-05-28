package main

type Talker interface {
	Talk()
}

type Human struct {
	name string
}

// human can talk
func (h Human) Talk() {
	println("My name is", h.name)
}

type Dog struct {
	name string
}

// dog can bark
func (d Dog) bark() {
	println("Woof!")
}

// we want to talk to dog, so we implement adapter for talking
type DogTalkerAdapter struct {
	Dog
}

// now dog can talk too
func (dta DogTalkerAdapter) Talk() {
	dta.bark()
}

func main() {
	h := Human{"Igor"}
	d := Dog{"Barsik"}
	dtAdapter := DogTalkerAdapter{Dog: d}
	talkers := []Talker{h, dtAdapter}
	for _, t := range talkers {
		t.Talk()
	}
}
