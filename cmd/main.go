package main

import (
	"fmt"

	"github.com/jroden2/stackforge/pkg/bundle"
)

func main() {
	fmt.Println("Mac setup tool")
	if bundles, err := bundle.LoadBundles("./bundles/"); err != nil {
		fmt.Println("Failed to load bundles", err)
		return
	} else {
		if len(bundles) == 0 {
			fmt.Println("No bundles found")
			return
		}

		fmt.Printf("Found %d bundles\n", len(bundles))
		for i, b := range bundles {
			fmt.Printf("%d %s\n", i+1, b.Title)
			fmt.Printf("  %s\n", b.Description)
		}
	}
}
