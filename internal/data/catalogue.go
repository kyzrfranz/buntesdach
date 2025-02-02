package data

import (
	"encoding/xml"
	"github.com/kyzrfranz/buntesdach/pkg/resources"
	"github.com/samber/lo"
)

type ItemsGetter[E resources.Entry] interface {
	GetItems() []E
}

type CatalogReader[C ItemsGetter[E], E resources.Entry] struct {
	fetcher CatalogFetcher
}

type CatalogFetcher interface {
	Fetch() ([]byte, error)
}

func NewCatalogReader[C ItemsGetter[E], E resources.Entry](fetcher CatalogFetcher) CatalogReader[C, E] {
	return CatalogReader[C, E]{
		fetcher: fetcher,
	}
}

func (r *CatalogReader[C, E]) GetEntry(id string) (*resources.Entry, error) {
	var e resources.Entry
	raw, _ := r.GetCatalogueEntry(id)
	e = *raw
	return &e, nil
}

func (r *CatalogReader[C, E]) GetCatalog() ([]E, error) {
	c, _ := r.readCatalog()
	return c.GetItems(), nil
}

func (r *CatalogReader[T, E]) GetCatalogueEntry(id string) (*E, error) {
	catalog, err := r.readCatalog()
	if err != nil {
		return nil, err
	}
	catalogEntry, ok := lo.Find(catalog.GetItems(), func(entry E) bool {
		var e resources.Entry
		e = entry
		return e.GetId() == id
	})
	if !ok {
		return nil, nil
	}

	return &catalogEntry, nil
}

func (r *CatalogReader[T, E]) readCatalog() (ItemsGetter[E], error) {
	data, err := r.fetcher.Fetch()
	if err != nil {
		return nil, err
	}

	var catalog T

	err = xml.Unmarshal(data, &catalog)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}
