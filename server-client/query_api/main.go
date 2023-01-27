package main

import (
	"fmt"
	"net/http"

	"query_api/service"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("router")
	router.HandleFunc("/posts", service.GetPosts).Methods("GET")
	router.HandleFunc("/postint", service.CreatePost).Methods("POST")
	router.HandleFunc("/post1", service.GetPost).Methods("GET")
	router.HandleFunc("/updatepost", service.UpdatePost).Methods("PUT")
	router.HandleFunc("/delete", service.DeletePost).Methods("DELETE")
	http.ListenAndServe(":8080", router)
	defer service.Db.Close()
}
