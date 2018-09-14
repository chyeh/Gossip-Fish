package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/olivere/elastic"
)

type datastore interface {
	searchArticles(*Query) *elastic.SearchResult
	searchComments(*Query) []SearchCommentsView
}

type elasticsearchClient struct {
	client *elastic.Client
	index  string
}

func newElasticsearchClient(elasticsearch string) *elasticsearchClient {
	c, err := elastic.NewClient(elastic.SetURL(elasticsearch))
	if err != nil {
		panic(err)
	}
	return &elasticsearchClient{
		client: c,
		index:  "gossip",
	}
}

func (c *elasticsearchClient) searchArticles(query *Query) *elastic.SearchResult {
	multiMatchQuery := elastic.
		NewMultiMatchQuery(query.Q,
			"article_title",
			"author",
			"ip",
			"content",
		).
		Type("phrase")
	result, err := c.client.Search().
		Index(c.index).
		Query(multiMatchQuery).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return result
}

func (c *elasticsearchClient) searchComments(query *Query) []SearchCommentsView {
	multiMatchQuery := elastic.
		NewMultiMatchQuery(query.Q,
			"messages.push_userid",
			"messages.push_ipdatetime",
			"messages.push_content",
		).
		Type("phrase")
	nestedQuery := elastic.
		NewNestedQuery("messages", multiMatchQuery).
		InnerHit(elastic.NewInnerHit())
	result, err := c.client.Search().
		Index(c.index).
		Query(nestedQuery).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	ans := make([]SearchCommentsView, 0)

	for _, hit := range result.Hits.Hits {
		articleModel := &ArticleModel{}
		loadModel(hit, articleModel)
		hitCommentModels := make([]*CommentModel, hit.InnerHits["messages"].Hits.TotalHits)
		for i, innerHit := range hit.InnerHits["messages"].Hits.Hits {
			hitCommentModel := &CommentModel{}
			loadModel(innerHit, hitCommentModel)
			hitCommentModels[i] = hitCommentModel
		}

		articleView := SearchCommentsView{
			Comments: make([]*CommentView, len(articleModel.Comments)),
			Hits:     make([]*CommentView, len(hitCommentModels)),
		}

		articleView.ID = articleModel.ID
		articleView.Title = articleModel.Title
		articleView.Author = articleModel.Author
		articleView.Board = articleModel.Board
		articleView.IP = articleModel.IP
		articleView.Time = parseANSICTime(articleModel.Time)
		articleView.Content = articleModel.Content
		for i, commentModel := range articleModel.Comments {
			commentView := &CommentView{}
			commentView.Account = commentModel.Account
			commentView.Message = commentModel.Message

			parsedIP, parsedTime := parseIPDateTime(articleView.Time.Year(), commentModel.IPDateTime)
			commentView.IP = parsedIP
			commentView.Time = parsedTime
			articleView.Comments[i] = commentView
		}

		for i, hitCommentModel := range hitCommentModels {
			hitCommentView := &CommentView{}
			hitCommentView.Account = hitCommentModel.Account
			hitCommentView.Message = hitCommentModel.Message

			parsedIP, parsedTime := parseIPDateTime(articleView.Time.Year(), hitCommentModel.IPDateTime)
			hitCommentView.IP = parsedIP
			hitCommentView.Time = parsedTime
			articleView.Hits[i] = hitCommentView
		}

		ans = append(ans, articleView)
	}
	return ans
}
