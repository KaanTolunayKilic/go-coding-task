package main

type Client struct {
}

type Post struct {
}

func NewClient() Client {
	return Client{}
}

func (c *Client) ListPosts(id int) ([]Post, error) {
	panic("Need to implement method!")
}
