package storage

import (
	"context"
	"elasticsearch-service/app/domain/models"
	"elasticsearch-service/app/infra/elastics"
	"elasticsearch-service/app/uerror"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/esutil"
)

type IProductElstic interface {
	IndexProduct(product models.ProductModel) (*esapi.Response, error)
	IndexProducts(product []models.ProductModel) (*esapi.Response, error)
}

var (
	defaultHeader = map[string]string{
		"Content-Type": "application/json",
	}
	ProductElastic IProductElstic
)

type productElastic struct{}

func init() {
	ProductElastic = &productElastic{}
}

func (p productElastic) IndexProduct(product models.ProductModel) (*esapi.Response, error) {
	buf, err := json.Marshal(product)
	if err != nil {
		log.Println("can't marshal product")
		return nil, uerror.InternalError(err, err.Error())
	}

	res, err := elastics.ElasticSearchClient().Index(
		product.Index(),
		esutil.NewJSONReader(buf),
		elastics.ElasticSearchClient().Index.WithDocumentID(strconv.Itoa(product.Id)),
		elastics.ElasticSearchClient().Index.WithRefresh("true"),
		elastics.ElasticSearchClient().Index.WithContext(context.Background()),
		elastics.ElasticSearchClient().Index.WithHeader(defaultHeader),
	)

	if err != nil {
		log.Println("cann't index product")
		return nil, uerror.InternalError(err, err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.IsError() {
		log.Println("response body index error")
		return nil, fmt.Errorf("response body index error %d", product.Id)
	}
	return nil, nil
}

func (p productElastic) IndexProducts(products []models.ProductModel) (*esapi.Response, error) {
	bi, err := esutil.NewBulkIndexer(models.BulkIndexConf())
	if err != nil {
		panic(err)
	}

	// Iterate over the products and add them to the bulk index request
	for _, product := range products {
		// Convert the product to a JSON string
		buf, err := json.Marshal(product)
		if err != nil {
			log.Println("cann't marshal product")
			return nil, uerror.InternalError(err, err.Error())
		}

		// Create a new bulk index request item
		item := esutil.BulkIndexerItem{
			Index:      product.Index(),
			DocumentID: strconv.Itoa(product.Id),
			Body:       esutil.NewJSONReader(buf),
		}

		// Add the item to the bulk index request
		err = bi.Add(context.Background(), item)
		if err != nil {
			// Handle error
		}
	}

	// Close the bulk index request
	err = bi.Close(context.Background())
	if err != nil {
		// Handle error
	}
	return nil, nil
}
