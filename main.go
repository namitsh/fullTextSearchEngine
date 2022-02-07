package main

import (
	"fmt"
)

var idx index

func main() {
	fmt.Println("Hello World")
	fmt.Println("Loading data from source")
	loadDocument("wikipedia.xml")
	//idx := make(index)
	fmt.Println("adding the dumps to create an inverted index")
	idx.add(dump.Documents)
	fmt.Println("Successfully created the inverted index")
	fmt.Println("Searching the documents")
	r := idx.search("Hello World")
	fmt.Println(len(r))
	for _, id := range r {
		fmt.Println(id)
	}
}
