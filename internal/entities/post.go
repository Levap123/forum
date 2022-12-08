package entities

type Post struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type User struct {
	Id       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}
