package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CharlieAlbert/deadlinks-checker/internal/checker"
	"github.com/CharlieAlbert/deadlinks-checker/internal/fetch"
	"github.com/CharlieAlbert/deadlinks-checker/internal/parser"
	"github.com/CharlieAlbert/deadlinks-checker/internal/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter webiste URL: ")
	input, _ := reader.ReadString('\n')
	url := strings.TrimSpace(input)

	htmlContent, err := fetch.FetchHTML(url)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		return
	}

	links, err := parser.ExtractLinks(htmlContent)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	normalized, err := utils.NormalizeLinks(links, url)
	if err != nil {
		fmt.Println("Error normalizing links:", err)
		return
	}

	fmt.Printf("\n Checking %d links...\n\n", len(normalized))

	statuses := checker.CheckLinksConcurrently(normalized)

	for _, result := range statuses {
		if result.Alive {
			fmt.Printf("✅ %s (%s)\n", result.URL, result.Status)
		} else {
			fmt.Printf("❌ %s (%s)\n", result.URL, result.Status)
		}
	}
}
