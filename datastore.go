package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/olivere/elastic"
)

type datastore interface {
	searchArticles(*Query) *elastic.SearchResult
	searchComments(*Query)
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
		NewMultiMatchQuery(query.Q, "article_title", "author", "ip", "content").
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

func (c *elasticsearchClient) searchComments(query *Query) {
	return
}
