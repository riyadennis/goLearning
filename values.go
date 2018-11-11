package main

import (
	"fmt"
)

type jumper interface {
	jump() string
	getName() string
}
type horse struct {
	name   string
	weight int
}
type gopher struct {
	name    string
	age     int
	isAdult bool
}

func (h *horse) jump() string {
	if h.weight > 2200 {
		return "cannot jump"
	}
	return "can jump"
}
func (g *gopher) jump() string {
	if g.age > 18 {
		return "can jump"
	}
	return "cannot jump"
}
func main() {
	for _, g := range getList() {
		fmt.Printf("jumping status of %s is  %s\n", g.getName(), g.jump())
	}
}
func getList() []jumper {
	phil := &gopher{"Phil", 19, true}
	bob := &gopher{"bob", 9, false}
	h := &horse{"barny", 234}
	return []jumper{phil, bob, h}
}
func (h *horse) getName() string {
	return h.name
}

func (g *gopher) getName() string {
	return g.name
}
