package main

import "strings"

// Post struct for post
type Post struct {
	ID     int
	UserID int
	Title  string
	Body   string
}

// Comment struct for comment
type Comment struct {
	ID     int
	PostID int
	Name   string
	Email  string
	Body   string
}

func (c *Comment) contains(substr string) bool {
	return strings.Contains(c.Body, substr) || strings.Contains(c.Email, substr)
}
