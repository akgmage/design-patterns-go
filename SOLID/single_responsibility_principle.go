package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
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


// breaks SRP

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.Stringify()), 0644)
	// 0644 readable by all the user groups, but writable by the user only.
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// best way to do it without breaking SRP
var lineSeparator = "\n"
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, lineSeparator)), 0644)
}


func main() {

}