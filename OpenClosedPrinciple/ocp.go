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

type Filter struct {

}

func (f * Filter) FilterByColor( products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}