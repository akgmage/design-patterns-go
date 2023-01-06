// Some objects are simple and can be created in a single constructor call
// Other objeccts require a lot of ceremony to create
// Having a factory function with 15 arguments is not productive
// Instead, opt for piecewise (piece-by-piece) construction
// BUilder provides an API for constructing an object step by step

// Builder : When piecewise object construction is complicated, provide an API for doing it in brief and clear manner

package main

import (
	"fmt"
	"strings"
)

type HtmlElement struct {
	name, text string
	elements []HtmlElement
}

func main () {
	text := "Hello"
	sb := strings.Builder{}
	sb.WriteString("<h1>")
	sb.WriteString(text)
	sb.WriteString("</h1>")
	fmt.Println(sb.String())

	words := []string{"hello", "go"}
	sb.Reset()
	// <ul><li>..</li><li>..</li></ul>
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())
}