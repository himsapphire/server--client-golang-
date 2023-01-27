package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPosts() {

	req, err := http.NewRequest("GET", "http://localhost:8080/posts", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s \n ", body)

	fmt.Printf("Response status : %s \n", resp.Status)

}

func CreatePost() {

	myJson := bytes.NewBuffer([]byte(`{
		"id" : "13",
		"name" : "hap",
		"balance" : "5500"
	}`))

	resp, err := c.Post("http://localhost:8080/postint", "application/json", myJson)

	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s \n ", body)

	fmt.Printf("Response status : %s \n", resp.Status)

}

func GetPost() {

	myJson1 := bytes.NewBuffer([]byte(`{
		"id" : "31"
	}`))

	req, err := http.NewRequest("GET", "http://localhost:8080/post1", myJson1)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s \n ", body)

	fmt.Printf("Response status : %s \n", resp.Status)

}

func UpdatePost() {

	myJson1 := bytes.NewBuffer([]byte(`{
		"id" : "30",
		"balance" : "9900"
	}`))

	req, err := http.NewRequest("PUT", "http://localhost:8080/updatepost", myJson1)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s \n ", body)

	fmt.Printf("Response status : %s \n", resp.Status)

}

func DeletePost() {

	myJson1 := bytes.NewBuffer([]byte(`{
		"id" : "32"
	}`))

	req, err := http.NewRequest("DELETE", "http://localhost:8080/delete", myJson1)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s \n ", body)

	fmt.Printf("Response status : %s \n", resp.Status)

}
