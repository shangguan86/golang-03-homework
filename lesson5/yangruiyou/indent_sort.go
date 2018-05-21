package main

import "fmt"

type Entries []Entry

type Entry struct {
	key      string
	value    string
	children Entries
}

func (entries Entries) Len() int {
	return len(entries)
}

func (entries Entries) Less(i, j int) bool {
	return entries[i].key < entries[j].key

}

func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}

func populateEntries(slice []string)Entries{
	
}

func SortedIndentedStrings(slice []string)[]string{
	entries :=popu
}


func main() {
	fmt.Println("|    Original    |    Sorted    |")
	fmt.Println("|----------------|--------------|")

}
