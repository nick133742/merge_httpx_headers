package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: merge_httpx_headers file1.txt file2.txt ...")
		os.Exit(1)
	}

	var mergedHeaders []string

	// Read each file and extract headers
	for _, filename := range os.Args[1:] {
		content, err := os.ReadFile(filename)
		if err != nil {
			continue
		}

		// Convert the content to string
		fileContent := string(content)

		// Find the index of the first empty line (header separator)
		separatorIndex := strings.Index(fileContent, "\r\n\r\n")
		// If the separator is not found, skip this file
		if separatorIndex == -1 {
			continue
		}
		// Extract and append headers to the mergedHeaders string
		headers := strings.Split(fileContent[:separatorIndex], "\r\n")
		mergedHeaders = append(mergedHeaders, headers...)

	}
	sort.Strings(mergedHeaders)
	sortedMergedHeaders := strings.Join(mergedHeaders, "\n")
	fmt.Println(sortedMergedHeaders)
}
