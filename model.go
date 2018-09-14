package main

import (
	"time"
)

type CommentModel struct {
	Account    string `json:"push_userid"`
	Message    string `json:"push_content"`
	IPDateTime string `json:"push_ipdatetime"`
}

type ArticleModel struct {
	ID       string         `json:"article_id"`
	Title    string         `json:"article_title"`
	Author   string         `json:"author"`
	Board    string         `json:"board"`
	IP       string         `json:"ip"`
	Time     string         `json:"date"`
	Content  string         `json:"content"`
	Comments []CommentModel `json:"messages"`
}

type CommentView struct {
	Account string    `json:"account"`
	Message string    `json:"message"`
	IP      string    `json:"ip"`
	Time    time.Time `json:"time"`
}

type SearchArticlesView struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Author   string        `json:"author"`
	Board    string        `json:"board"`
	IP       string        `json:"ip"`
	Time     time.Time     `json:"time"`
	Content  string        `json:"content"`
	Comments []CommentView `json:"comments"`
}

type SearchCommentsView struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Author   string        `json:"author"`
	Board    string        `json:"board"`
	IP       string        `json:"ip"`
	Time     time.Time     `json:"time"`
	Content  string        `json:"content"`
	Comments []CommentView `json:"comments"`
	Hits     []CommentView `json:"hits"`
}

type Query struct {
	Q string `form:"q"`
}
