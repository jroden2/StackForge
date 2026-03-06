package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jroden2/stackforge/pkg/bundle"
	"github.com/jroden2/stackforge/pkg/installer"
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

		fmt.Println("Choose a bundle:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		selection, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid selection")
			return
		}

		index := selection - 1
		if index < 0 || index >= len(bundles) {
			fmt.Println("Invalid selection")
			return
		}

		err = installer.InstallBundle(bundles[index])
		if err != nil {
			fmt.Println("Installation failed:", err)
		}
	}
}
