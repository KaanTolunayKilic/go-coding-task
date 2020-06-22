package main

import (
	"fmt"
	"strings"
)

// PrintPost prints post on console
func PrintPost(post Post) {
	fmt.Printf("Post: %s (userId=%d)\n", post.Title, post.UserID)
	fmt.Println(" >", removeLineBreaks(post.Body))
}

// PrintComment prints comment on console
func PrintComment(comment Comment) {
	fmt.Printf("    Comment: %s (by %s)\n", removeLineBreaks(comment.Body), comment.Email)
}

// PrintDivider prints dividing line on console
func PrintDivider() {
	fmt.Println("-----------------------------------------------------------------------")
}

func removeLineBreaks(text string) string {
	return strings.Replace(text, "\n", "", -1)
}
