package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const apiURL string = "https://jsonplaceholder.typicode.com"

type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

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

func NewClient(httpClient *http.Client) Client {
	url, _ := url.Parse(apiURL)
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return Client{httpClient: httpClient, BaseURL: url}
}

func (c *Client) UserPosts(userId int) ([]Post, error) {
	params := map[string]string{
		"userId": fmt.Sprintf("%d", userId),
	}

	req, err := c.newRequest(http.MethodGet, "/posts", params)
	if err != nil {
		return nil, err
	}

	var posts []Post
	_, err = c.do(req, &posts)
	return posts, err
}

func (c *Client) PostComments(postId int) ([]Comment, error) {
	params := map[string]string{
		"postId": fmt.Sprintf("%d", postId),
	}

	req, err := c.newRequest(http.MethodGet, "/comments", params)
	if err != nil {
		return nil, err
	}

	var comments []Comment
	_, err = c.do(req, &comments)
	return comments, err
}

func (c *Client) newRequest(method, path string, params map[string]string) (*http.Request, error) {
	relPath := &url.URL{Path: path}
	url := c.BaseURL.ResolveReference(relPath).String()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	query := req.URL.Query()
	for key, val := range params {
		query.Add(key, val)
	}
	req.URL.RawQuery = query.Encode()

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
