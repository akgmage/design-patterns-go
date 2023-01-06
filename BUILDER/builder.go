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

const (
	indentSize = 2
)
type HtmlElement struct {
	name, text string
	elements []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize * indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName, HtmlElement{rootName, "", []HtmlElement{}}}
}

// have string representaion for builder 
func (b *HtmlBuilder) String() string {
	return b.root.String()
}

// method for adding a child
func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
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

	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "go")
	b.AddChild("li", "haha")
	fmt.Println(b.String())
}