package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	userIDPtr := flag.Int("userId", -1, "User id")
	filterPtr := flag.String("filter", "", "Filter comments")
	flag.Parse()

	client := NewClient(nil)

	posts, err := client.UserPosts(*userIDPtr)
	if err != nil {
		log.Fatalf("List posts failed with error: %v", err)
	}

	fmt.Printf("Posts and comments for user with id %d:\n", *userIDPtr)
	for _, post := range posts {
		PrintPost(post)
		comments, _ := client.PostComments(post.ID)
		for _, comment := range comments {
			if strings.Contains(comment.Body, *filterPtr) {
				PrintComment(comment)
			}
		}
		PrintDivider()
	}
}
