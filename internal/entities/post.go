package entities

type Post struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	UserId   int    `json:"userId,omitempty"`
	Likes    int
	Dislikes int
}

type Action struct {
	Id     int
	Vote   int
	UserId int
	PostId int
}
