package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {

	fmt.Println("enter into getposts")

	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	result, err := Db.Query("SELECT id, name, balance from bal")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Name, &post.Balance)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	//response := jsonresponse{Msg: "Success", Data: posts}
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("entry to insertfunc")

	stmt, err := Db.Prepare("INSERT INTO bal(id,name,balance) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]
	fmt.Println(id)
	id1, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println("id is", id1)
	name := keyVal["name"]
	balance := keyVal["balance"]
	balance1, _ := strconv.ParseInt(balance, 10, 64)

	_, err = stmt.Exec(id1, name, balance1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "New post was created")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter into getpost")
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	id := keyVal["id"]
	id1, _ := strconv.ParseInt(id, 10, 64)
	result, err := Db.Query("SELECT * FROM bal WHERE id = ?", id)
	if err != nil {

		response := jsonresponse{Msg: "Failure"}
		json.NewEncoder(w).Encode(response)

		fmt.Println(err)

	}
	defer result.Close()
	var post Post
	for result.Next() {
		err := result.Scan(&post.ID, &post.Name, &post.Balance)
		if err != nil {
			fmt.Println(err)
		}
		if post.ID == int(id1) {
			break
		}
	}
	if(post.ID== int(id1)) {
		response := jsonresponse{Msg: "Success", Data: post}
			json.NewEncoder(w).Encode(response)
	}else {
		response2 := jsonresponse{Msg: "Failure", Data: post}
		json.NewEncoder(w).Encode(response2)
	}
	
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	fmt.Println("enter into updatepost")
	w.Header().Set("Content-Type", "application/json")
	

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newbalance := keyVal["balance"]
	id := keyVal["id"]

	id1, _ := strconv.ParseInt(id, 10, 64)
	newbalance1, _ := strconv.ParseInt(newbalance, 10, 64)


	stmt, err := Db.Query("UPDATE bal SET balance = ? WHERE id = ?", newbalance1, id1)
	if err != nil {
		response := jsonresponse{Msg: "id not found"}
		json.NewEncoder(w).Encode(response)
		fmt.Println(err)
	}else{
		fmt.Println(stmt)
		fmt.Fprintf(w, "Post with ID = %v was updated", id)
	}
	//_, err = stmt.Exec(newbalance, params["id"])

	
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter into deletepost")
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]
	st, err := Db.Query("DELETE FROM bal WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(st)

	fmt.Fprintf(w, "Post with ID = %s was deleted", id)
}
