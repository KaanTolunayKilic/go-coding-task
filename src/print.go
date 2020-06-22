package main

import (
	"fmt"
	"strings"
)

func PrintPost(post Post) {
	fmt.Printf("Post: %s (userId=%d)\n", post.Title, post.UserID)
	fmt.Println(" >", removeLineBreaks(post.Body))
}

func PrintComment(comment Comment) {
	fmt.Printf("    Comment: %s (by %s)\n", removeLineBreaks(comment.Body), comment.Email)
}

func PrintDivider() {
	fmt.Println("===========================================================")
}

func removeLineBreaks(text string) string {
	return strings.Replace(text, "\n", "", -1)
}
