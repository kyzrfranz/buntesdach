package data

import (
	"encoding/json"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"log/slog"
)

type PoliticianCatalogReader struct {
	data *v1.PoliticianCatalog
}

type PoliticianFetcher interface {
	Fetch() ([]byte, error)
}

func NewPoliticianCatalogReader(fetcher PoliticianFetcher) (*PoliticianCatalogReader, error) {

	data, err := fetcher.Fetch()
	if err != nil {
		return nil, err
	}

	var politicianCatalog v1.PoliticianCatalog

	err = json.Unmarshal(data, &politicianCatalog)
	if err != nil {
		return nil, err
	}

	return &PoliticianCatalogReader{
		data: &politicianCatalog,
	}, nil
}

func (r *PoliticianCatalogReader) GetPoliticianCatalog() *v1.PoliticianCatalog {
	return r.data
}

func (r *PoliticianCatalogReader) GetPolitician(id string) *v1.Politician {
	if r.data == nil {
		slog.Error("PoliticianCatalogReader.GetPolitician: data is nil")
		return nil
	}
	catalog := *r.data
	catalogEntry, ok := catalog[id]
	if !ok {
		return nil
	}
	return &catalogEntry
}
