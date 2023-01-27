package service

type Post struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type jsonresponse struct {
	Msg  string
	Data Post
}
