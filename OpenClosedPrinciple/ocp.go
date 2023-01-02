package main

type Color int 

const (
	red Color = iota
	green
	bluered
)
type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name string
	color Color
	size Size
}