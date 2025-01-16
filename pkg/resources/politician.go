package resources

import (
	"context"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
)

type politicianHandler struct {
	getter CatalogueGetter[v1.PersonListEntry]
}

func NewPoliticianHandler(getter CatalogueGetter[v1.PersonListEntry]) Handler[v1.Politician] {
	return politicianHandler{
		getter: getter,
	}
}

func (p politicianHandler) List(ctx context.Context) []v1.Politician {
	//TODO implement me
	panic("implement me")
}

func (p politicianHandler) Get(ctx context.Context, id string) (*v1.Politician, error) {

	//catalogEntry, err := p.getter.GetCatalogueEntry(id)
	//if err != nil {
	//	return nil, err
	//}
	//
	//bUrl, err := url.Parse(catalogEntry.InfoXMLURL)
	//if err != nil {
	//	return nil, err
	//}
	//data, err := myhttp.FetchUrl(bUrl)
	//
	//var bio v1.Politician
	//if err = xml.Unmarshal(data, &bio); err != nil {
	//	return nil, err
	//}
	//
	//return &bio, nil
	return nil, nil
}

func (p politicianHandler) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p politicianHandler) Create(ctx context.Context, item *v1.Politician) (*v1.Politician, error) {
	//TODO implement me
	panic("implement me")
}

func (p politicianHandler) Update(ctx context.Context, oldItem *v1.Politician, newItem *v1.Politician) (*v1.Politician, error) {
	//TODO implement me
	panic("implement me")
}

func (p politicianHandler) Name() string {
	//TODO implement me
	panic("implement me")
}
