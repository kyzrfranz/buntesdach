package resources

import (
	"context"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
)

type catalogueHandler struct {
	getter CatalogueGetter
}

type CatalogueGetter interface {
	GetPoliticianCatalog() (*v1.PersonCatalog, error)
	GetCatalogueEntry(id string) (*v1.PersonListEntry, error)
}

func NewCatalogueHandler(getter CatalogueGetter) Handler[v1.PersonListEntry] {
	return catalogueHandler{
		getter: getter,
	}
}

func (s catalogueHandler) List(ctx context.Context) []v1.PersonListEntry {
	catalog, err := s.getter.GetPoliticianCatalog()
	if err != nil {
		return nil
	}
	return catalog.Persons
}

func (s catalogueHandler) Get(ctx context.Context, id string) (*v1.PersonListEntry, error) {
	return s.getter.GetCatalogueEntry(id)
}

func (s catalogueHandler) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s catalogueHandler) Create(ctx context.Context, item *v1.PersonListEntry) (*v1.PersonListEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (s catalogueHandler) Update(ctx context.Context, oldItem *v1.PersonListEntry, newItem *v1.PersonListEntry) (*v1.PersonListEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (s catalogueHandler) Name() string {
	//TODO implement me
	panic("implement me")
}
