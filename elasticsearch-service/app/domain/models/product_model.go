package models

import (
	"time"

	"github.com/elastic/go-elasticsearch/esutil"
)

type ProductModel struct {
	Id        int
	Name      string
	CreatedAt string
	UpdatedAt string
}

func (*ProductModel) Index() string {
	return "product_index"
}

func BulkIndexConf() esutil.BulkIndexerConfig {
	return esutil.BulkIndexerConfig{
		Index:         "product_index",
		NumWorkers:    1,
		FlushBytes:    1024 * 1024,
		FlushInterval: 30 * time.Second,
	}
}
