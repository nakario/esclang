package esc

type module struct {
	name         rune
	labels       map[rune]int
	instructions []Instruction
}
