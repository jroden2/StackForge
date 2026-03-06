package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jroden2/stackforge/pkg/bundle"
	"github.com/jroden2/stackforge/pkg/installer"

	_ "github.com/jroden2/stackforge/pkg/installer/managers"
)

func main() {
	fmt.Println("Mac setup tool")

	bundles, err := bundle.LoadBundles("./bundles/")
	if err != nil {
		fmt.Println("Failed to load bundles:", err)
		os.Exit(1)
	}

	if len(bundles) == 0 {
		fmt.Println("No bundles found")
		os.Exit(1)
	}

	fmt.Printf("Found %d bundles\n\n", len(bundles))
	for i, b := range bundles {
		fmt.Printf("  %d) %s - %s\n", i+1, b.Title, b.Description)
	}

	fmt.Print("\nChoose a bundle (0 to exit): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		os.Exit(1)
	}

	selection, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || selection < 0 || selection > len(bundles) {
		fmt.Println("Invalid selection")
		os.Exit(1)
	}

	if selection == 0 {
		fmt.Println("Exiting")
		os.Exit(0)
	}

	if err := installer.InstallBundle(bundles[selection-1]); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}
}
