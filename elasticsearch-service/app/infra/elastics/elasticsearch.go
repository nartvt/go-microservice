package elastics

import (
	"elasticsearch-service/app/config"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch"
)

var elasticClient *elasticsearch.Client

func init() {
	load()
}
func load() {
	conf := config.Config
	address := fmt.Sprintf("%s:%d", conf.ElasticSearch.Host, conf.ElasticSearch.Port)
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
		Username: conf.ElasticSearch.UserName,
		Password: conf.ElasticSearch.Password,
	}

	elasticClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("Could not create client elastic ", err.Error())
		panic(err)
	}
	if elasticClient.Ping == nil {
		err := fmt.Errorf("could not ping to elastic %s", address)
		panic(err)
	}
}

func ElasticSearchClient() *elasticsearch.Client {
	if elasticClient == nil {
		load()
	}
	return elasticClient
}
