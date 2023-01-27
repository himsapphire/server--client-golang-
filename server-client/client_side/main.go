package main

import (
	//"bytes"

	"net/http"
	"time"
)

var c http.Client

func main() {
	c = http.Client{Timeout: time.Duration(10) * time.Second}
	//GetPosts()
	//CreatePost()
	//GetPost()
	//UpdatePost()
	DeletePost()
}
