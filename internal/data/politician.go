package data

import (
	"encoding/xml"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/samber/lo"
)

type PoliticianCatalogReader struct {
	data    *v1.PersonCatalog
	fetcher PoliticiansFetcher
}

type PoliticiansFetcher interface {
	Fetch() ([]byte, error)
}

func NewPoliticianCatalogReader(fetcher PoliticiansFetcher) (*PoliticianCatalogReader, error) {

	return &PoliticianCatalogReader{
		fetcher: fetcher,
	}, nil
}

func (r *PoliticianCatalogReader) GetPoliticianCatalog() (*v1.PersonCatalog, error) {
	return r.readCatalog()
}

func (r *PoliticianCatalogReader) GetCatalogueEntry(id string) (*v1.PersonListEntry, error) {
	catalog, err := r.readCatalog()
	if err != nil {
		return nil, err
	}
	catalogEntry, ok := lo.Find(catalog.Persons, func(entry v1.PersonListEntry) bool {
		return entry.Id.Value == id
	})
	if !ok {
		return nil, nil
	}

	return &catalogEntry, nil
}

func (r *PoliticianCatalogReader) readCatalog() (*v1.PersonCatalog, error) {
	data, err := r.fetcher.Fetch()
	if err != nil {
		return nil, err
	}

	var politicianCatalog v1.PersonCatalog

	err = xml.Unmarshal(data, &politicianCatalog)
	if err != nil {
		return nil, err
	}

	return &politicianCatalog, nil
}
