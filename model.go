package main

type Comment struct {
	IPDateTime  string `json:"push_ipdatetime"`
	UserAccount string `json:"push_userid"`
	Message     string `json:"push_content"`
}

type Article struct {
	ID       string    `json:"article_id"`
	Title    string    `json:"article_title"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
	IP       string    `json:"ip"`
	Comments []Comment `json:"messages"`
}

type Query struct {
	Q string `form:"q"`
}
