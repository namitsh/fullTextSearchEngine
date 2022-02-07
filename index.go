package main

import (
	"fmt"
	"regexp"
)

type index map[string][]int

func (idx index) add(docs []Document) {
	for _, doc := range docs {
		for _, token := range tokenize(doc.Text) {
			ids := idx[token]
			// TODO : need a better way for this.
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func (idx index) search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}

func searchDocument(term string) {
	var r []string
	docs := dump.Documents
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)
	for _, doc := range docs {
		if re.MatchString(doc.Text) {
			r = append(r, doc.Title)
		}
	}
	fmt.Println(r)
}
