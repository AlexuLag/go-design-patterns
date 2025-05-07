package main

import "fmt"

type Research struct {
	// relationships Relationships
	browser RelationshipBrowser // low-level
}

func (r *Research) Investigate() {

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}
