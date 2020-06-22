package main

import "flag"

func main() {
	userIDPtr := flag.Int("userId", -1, "User id")
	filterPtr := flag.String("filter", "", "Filter comments")
	flag.Parse()

	client := NewClient()
	client.ListPosts(*userIDPtr)

	if *filterPtr != "" {
		// Filter post
	}
}
