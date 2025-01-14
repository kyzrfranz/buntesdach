package resources

import (
	"context"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
)

type politicianHandler struct {
	getter PoliticianGetter
}

type PoliticianGetter interface {
	GetPoliticianCatalog() *v1.PoliticianCatalog
}

func NewPoliticianHandler(getter PoliticianGetter) Handler[v1.Politician] {
	return politicianHandler{
		getter: getter,
	}
}

func (s politicianHandler) List(ctx context.Context) []v1.Politician {
	catalog := s.getter.GetPoliticianCatalog()
	var politicians []v1.Politician
	for _, politician := range *catalog {
		politicians = append(politicians, politician)
	}
	return politicians
}

func (s politicianHandler) Get(ctx context.Context, id string) (*v1.Politician, error) {
	catalog := s.getter.GetPoliticianCatalog()
	politician, ok := (*catalog)[id]
	if !ok {
		return nil, nil
	}
	return &politician, nil
}

func (s politicianHandler) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s politicianHandler) Create(ctx context.Context, item *v1.Politician) (*v1.Politician, error) {
	//TODO implement me
	panic("implement me")
}

func (s politicianHandler) Update(ctx context.Context, oldItem *v1.Politician, newItem *v1.Politician) (*v1.Politician, error) {
	//TODO implement me
	panic("implement me")
}

func (s politicianHandler) Name() string {
	//TODO implement me
	panic("implement me")
}
