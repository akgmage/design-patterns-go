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


// breaks srp

func (j *Journal) save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.Stringify()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

func main() {

}