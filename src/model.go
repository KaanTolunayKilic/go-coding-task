package main

import "strings"

type Post struct {
	ID     int
	UserID int
	Title  string
	Body   string
}

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
