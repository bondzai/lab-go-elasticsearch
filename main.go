package main

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	logger.SetLevel(logrus.InfoLevel)

	esConfig := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		logger.Fatal(err)
	}

	// Log some events
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// Index log events in Elasticsearch
	indexName := "logs"
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: "1",
		Body:       strings.NewReader(`{"message": "This is an informational message"}`),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		logger.Fatal(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		logger.Fatalf("Failed to index log event: %s", res.Status())
	} else {
		logger.Info("Log event indexed successfully")
	}
}
