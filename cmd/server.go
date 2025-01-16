package main

import (
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/kyzrfranz/buntesdach/internal/data"
	"github.com/kyzrfranz/buntesdach/internal/http"
	"github.com/kyzrfranz/buntesdach/internal/rest"
	"github.com/kyzrfranz/buntesdach/internal/upstream"
	"github.com/kyzrfranz/buntesdach/pkg/resources"
	"log/slog"
	"net/url"
	"os"
)

func main() {

	dataUrl, err := url.Parse("https://www.bundestag.de/xml/v2/mdb/index.xml") // TODO config
	if err != nil {
		bail("parse data url", err)
	}

	listFetcher := &upstream.XMLFetcher{Url: dataUrl}
	pReader := data.NewCatalogReader[v1.PersonCatalog, v1.PersonListEntry](listFetcher)
	if err != nil {
		bail("create politician catalog reader", err)
	}

	apiServer := http.NewApiServer(8080)

	apiServer.Use(http.MiddlewareRecovery)
	apiServer.Use(http.MiddlewareCORS)

	pCatalogHandler := rest.NewHandler[v1.PersonListEntry](resources.NewCatalogueHandler[v1.PersonListEntry](&pReader))
	apiServer.AddHandler("/politicians", pCatalogHandler.List)
	apiServer.AddHandler("/politicians/{id}", pCatalogHandler.Get)

	pHandler := resources.NewDetailHandler[v1.Politician](&pReader)
	ppHandler := rest.NewHandler[v1.Politician](pHandler)
	apiServer.AddHandler("/politicians/{id}/bio", ppHandler.Get)

	committeeUrl, err := url.Parse("https://www.bundestag.de/xml/v2/ausschuesse/index.xml") // TODO config
	if err != nil {
		bail("parse data url", err)
	}
	committeeFetcher := &upstream.XMLFetcher{Url: committeeUrl}
	cReader := data.NewCatalogReader[v1.CommitteeCatalog, v1.CommitteeListEntry](committeeFetcher)
	pCommitteeHandler := rest.NewHandler[v1.CommitteeListEntry](resources.NewCatalogueHandler[v1.CommitteeListEntry](&cReader))

	apiServer.AddHandler("/committees", pCommitteeHandler.List)
	apiServer.AddHandler("/committees/{id}", pCommitteeHandler.Get)

	cDetailHandler := resources.NewDetailHandler[v1.CommitteeDetails](&cReader)
	cHandler := rest.NewHandler[v1.CommitteeDetails](cDetailHandler)

	apiServer.AddHandler("/committees/{id}/detail", cHandler.Get)

	apiServer.ListenAndServe()
}

func bail(stage string, err error) {
	slog.Error("server bailing out", slog.String("stage", stage), "error", err)
	os.Exit(1)
}
