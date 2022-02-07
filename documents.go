package main

import (
	"encoding/xml"
	"os"
)

type DocumentList struct {
	Documents []Document `xml:"doc"`
}

var dump DocumentList

// Document create a document struct
type Document struct {
	ID    int
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Text  string `xml:"abstract"`
}

func loadDocument(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = xml.NewDecoder(f).Decode(&dump)
	if err != nil {
		panic(err)
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
}
