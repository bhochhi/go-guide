package main

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func removeTagsWithAttributes(input, attributeName string) (string, error) {
	tokenizer := html.NewTokenizer(bytes.NewReader([]byte(input)))
	var output string
	var inTagToRemove bool

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return output, nil // End of the document
		case html.TextToken:
			// If we're not in a tag to remove, append the text to the output
			if !inTagToRemove {
				output += string(tokenizer.Text())
			}
		case html.StartTagToken:
			token := tokenizer.Token()
			// Check if the tag has the specified attribute
			if _, exists := getAttribute(token, attributeName); exists {
				// If it does, set the flag to true to indicate we're in a tag to remove
				inTagToRemove = true
			} else {
				// If it doesn't, append the entire start tag to the output
				output += token.String()
			}
		case html.EndTagToken:
			// If we're in a tag to remove, set the flag back to false
			if inTagToRemove {
				inTagToRemove = false
			}
		}
	}

	return output, nil
}

func getAttribute(t html.Token, attributeName string) (string, bool) {
	for _, attr := range t.Attr {
		if attr.Key == attributeName {
			return attr.Val, true
		}
	}
	return "", false
}

func extractTextFromHTML(htmlString string) (string, error) {
	var result string

	reader := strings.NewReader(htmlString)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return result, nil // End of the document
		case html.TextToken:
			// Append text content to the result
			result += strings.TrimSpace(string(tokenizer.Text()))
		}
	}

	return result, nil
}

func main() {
	sourceHTML := `<div class="keep">This should be kept.</div>
				   <div remove-attr="remove">This and its content should be removed.</div>
				   <div></div>
				   <div>This should be kept too.</div>`

	result, err := removeTagsWithAttributes(sourceHTML, "remove-attr")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resultWithoutEmptyTags, err := extractTextFromHTML(result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(resultWithoutEmptyTags)
}
