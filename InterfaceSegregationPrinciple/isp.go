package main

type Document struct {

}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

func main() {

}