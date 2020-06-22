package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	userIDPtr := flag.Int("userId", -1, "User id *required")
	filterPtr := flag.String("filter", "", "Filter comments")
	flag.Parse()

	if *userIDPtr == -1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := NewClient(nil)
	fmt.Printf("Posts and comments for user with id %d:\n", *userIDPtr)
	for _, post := range client.posts(*userIDPtr) {
		PrintPost(post)
		for _, comment := range client.comments(&post) {
			if comment.contains(*filterPtr) {
				PrintComment(comment)
			}
		}
		PrintDivider()
	}
}

func (c *Client) comments(p *Post) []Comment {
	comments, err := c.PostComments(p.ID)
	if err != nil {
		log.Fatalf("Could not request comments for post with id %d: %s", p.ID, err)
	}
	return comments
}

func (c *Client) posts(userID int) []Post {
	posts, err := c.UserPosts(userID)
	if err != nil {
		log.Fatalf("Could not request posts for user with id %d: %s", userID, err)
	}
	return posts
}
