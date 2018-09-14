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
	ID       string          `json:"article_id"`
	Title    string          `json:"article_title"`
	Author   string          `json:"author"`
	Board    string          `json:"board"`
	IP       string          `json:"ip"`
	Time     string          `json:"date"`
	Content  string          `json:"content"`
	Comments []*CommentModel `json:"messages"`
}

type CommentView struct {
	Account string    `json:"account"`
	Message string    `json:"message"`
	IP      string    `json:"ip"`
	Time    time.Time `json:"time"`
}

func newCommentView(year int, commentModel *CommentModel) *CommentView {
	parsedIP, parsedTime := parseIPDateTime(year, commentModel.IPDateTime)
	return &CommentView{
		Account: commentModel.Account,
		Message: commentModel.Message,
		IP:      parsedIP,
		Time:    parsedTime,
	}
}

type SearchArticlesView struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Author   string         `json:"author"`
	Board    string         `json:"board"`
	IP       string         `json:"ip"`
	Time     time.Time      `json:"time"`
	Content  string         `json:"content"`
	Comments []*CommentView `json:"comments"`
}

type SearchCommentsView struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Author   string         `json:"author"`
	Board    string         `json:"board"`
	IP       string         `json:"ip"`
	Time     time.Time      `json:"time"`
	Content  string         `json:"content"`
	Comments []*CommentView `json:"comments"`
	Hits     []*CommentView `json:"hits"`
}

func newSearchCommentsView(articleModel *ArticleModel, hitCommentModels []*CommentModel) *SearchCommentsView {
	articleView := &SearchCommentsView{
		ID:       articleModel.ID,
		Title:    articleModel.Title,
		Author:   articleModel.Author,
		Board:    articleModel.Board,
		IP:       articleModel.IP,
		Time:     parseANSICTime(articleModel.Time),
		Content:  articleModel.Content,
		Comments: make([]*CommentView, len(articleModel.Comments)),
		Hits:     make([]*CommentView, len(hitCommentModels)),
	}
	for i, commentModel := range articleModel.Comments {
		articleView.Comments[i] = newCommentView(articleView.Time.Year(), commentModel)
	}
	for i, hitCommentModel := range hitCommentModels {
		articleView.Hits[i] = newCommentView(articleView.Time.Year(), hitCommentModel)
	}
	return articleView
}

type Query struct {
	Q string `form:"q"`
}
