package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/olivere/elastic"
)

type datastore interface {
	searchArticles(*Query) *elastic.SearchResult
	searchComments(*Query) []*SearchCommentsView
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

func (c *elasticsearchClient) searchComments(query *Query) []*SearchCommentsView {
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
	searchResult, err := c.client.Search().
		Index(c.index).
		Query(nestedQuery).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	result := make([]*SearchCommentsView, searchResult.TotalHits())
	for i, hit := range searchResult.Hits.Hits {
		articleModel := &ArticleModel{}
		loadModel(hit, articleModel)
		hitCommentModels := make([]*CommentModel, hit.InnerHits["messages"].Hits.TotalHits)
		for i, innerHit := range hit.InnerHits["messages"].Hits.Hits {
			hitCommentModel := &CommentModel{}
			loadModel(innerHit, hitCommentModel)
			hitCommentModels[i] = hitCommentModel
		}

		result[i] = newSearchCommentsView(articleModel, hitCommentModels)
	}
	return result
}
