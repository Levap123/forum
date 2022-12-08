package entities

type Post struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

