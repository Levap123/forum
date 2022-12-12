package entities

type Post struct {
	Id      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Body    string `json:"body,omitempty"`
	Actions int    `json:"actions,omitempty"`
	UserId  int    `json:"userId,omitempty"`
}
