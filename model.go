package main

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

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
	Metadata *Metadata            `json:"_metadata"`
	Records  []*SearchArticleView `json:"records"`
}

type SearchArticleView struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Author   string         `json:"author"`
	Board    string         `json:"board"`
	IP       string         `json:"ip"`
	Time     time.Time      `json:"time"`
	Content  string         `json:"content"`
	Comments []*CommentView `json:"comments"`
}

func newSearchArticleView(articleModel *ArticleModel) *SearchArticleView {
	view := &SearchArticleView{
		ID:       articleModel.ID,
		Title:    articleModel.Title,
		Author:   articleModel.Author,
		Board:    articleModel.Board,
		IP:       articleModel.IP,
		Time:     parseANSICTime(articleModel.Time),
		Content:  articleModel.Content,
		Comments: make([]*CommentView, len(articleModel.Comments)),
	}
	for i, commentModel := range articleModel.Comments {
		view.Comments[i] = newCommentView(view.Time.Year(), commentModel)
	}
	return view
}

type SearchCommentsView struct {
	Metadata *Metadata            `json:"_metadata"`
	Records  []*SearchCommentView `json:"records"`
}

type SearchCommentView struct {
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

func newSearchCommentView(articleModel *ArticleModel, hitCommentModels []*CommentModel) *SearchCommentView {
	view := &SearchCommentView{
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
		view.Comments[i] = newCommentView(view.Time.Year(), commentModel)
	}
	for i, hitCommentModel := range hitCommentModels {
		view.Hits[i] = newCommentView(view.Time.Year(), hitCommentModel)
	}
	return view
}

type Query struct {
	Q      string `form:"q"`
	Cursor int    `form:"cursor" validate:"min=0"`
	Limit  int    `form:"limit" validate:"min=0"`
}

func newQuery() *Query {
	return &Query{
		Q:      "",
		Cursor: 0,
		Limit:  10,
	}
}

type Metadata struct {
	Cursor     int `json:"cursor"`
	Limit      int `json:"limit"`
	TotalCount int `json:"total_count"`
	NextCursor int `json:"next_cursor"`
}

func newMetadata(cursor, limit, totalCount int) *Metadata {
	meta := &Metadata{
		Cursor:     cursor,
		Limit:      limit,
		TotalCount: totalCount,
		NextCursor: cursor + limit,
	}
	if cursor+limit >= totalCount {
		meta.NextCursor = -1
	}
	return meta
}
