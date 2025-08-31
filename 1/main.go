package main

import "fmt"

type Human struct{}

func (h *Human) Run() {
	fmt.Println("Running")
}

func (h *Human) Say() {
	fmt.Println("Saying")
}

type Action struct {
	Human
}

func main() {
	h := Human{}
	a := Action{}

	h.Run()
	h.Say()

	a.Run()
	a.Say()
}
