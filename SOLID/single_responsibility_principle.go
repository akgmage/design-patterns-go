package main

import (
	"fmt"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	entryCount--
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

func (j *Journal) Stringify() string {
	return strings.Join(j.entries, "\n")
}


func main() {

}