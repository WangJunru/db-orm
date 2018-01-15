package model

import (
	"net/http"
	"sync"
	"time"

	elastic "gopkg.in/olivere/elastic.v2"
)

const (
	esClientDefaultMaxRetires  = 3
	esClientDefaultHTTPTimeout = 10 * time.Second
)

var (
	_es_cfg         ESConfig
	_es_client      *ESClient
	_es_client_once sync.Once
)

func ElasticSetup(cfg ESConfig) {
	_es_cfg = cfg
}

func ElasticClient() *ESClient {
	_es_client_once.Do(func() {
		cli, err := _es_cfg.NewClient()
		if err != nil {
			panic(err)
		}

		_es_client = cli
	})

	return _es_client
}

type ESConfig struct {
	Endpoints   []string
	MaxRetries  int
	EnableGzip  bool
	HTTPTimeout time.Duration
	IndexName   string
}

func (e *ESConfig) NewClient() (*ESClient, error) {
	endpoints, maxRetries, enableGzip, httpTimeout := e.Endpoints, e.MaxRetries, e.EnableGzip, e.HTTPTimeout

	if maxRetries <= 0 {
		maxRetries = esClientDefaultMaxRetires
	}

	if httpTimeout <= 0 {
		httpTimeout = esClientDefaultHTTPTimeout
	}

	cli, err := elastic.NewClient(
		elastic.SetURL(endpoints...),
		elastic.SetMaxRetries(maxRetries),
		elastic.SetGzip(enableGzip),
		elastic.SetHttpClient(&http.Client{Timeout: httpTimeout}),
	)

	if err != nil {
		return nil, err
	}

	return &ESClient{
		Client:    cli,
		IndexName: e.IndexName,
	}, nil
}

type ESClient struct {
	*elastic.Client
	IndexName string
}

func (e *ESClient) IndexService(index string) *elastic.IndexService {
	if index == "" {
		index = e.IndexName
	}

	return e.Client.Index().Index(index)
}

func (e *ESClient) PutMappingService(index string) *elastic.PutMappingService {
	if index == "" {
		index = e.IndexName
	}

	return e.Client.PutMapping().Index(index)
}
