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

	dataUrl, err := url.Parse("https://www.bundestag.de/static/appdata/sitzplan/data.json") // TODO config
	if err != nil {
		bail("parse data url", err)
	}

	fetcher := &upstream.Politicians{Url: dataUrl}
	pReader, err := data.NewPoliticianCatalogReader(fetcher)
	if err != nil {
		bail("create politician catalog reader", err)
	}

	apiServer := http.NewApiServer(8001)

	apiServer.Use(http.MiddlewareRecovery)
	apiServer.Use(http.MiddlewareCORS)

	pCatalogHandler := rest.NewHandler[v1.Politician](resources.NewPoliticianHandler(pReader))
	apiServer.AddHandler("/politicians", pCatalogHandler.List)
	apiServer.AddHandler("/politicians/{id}", pCatalogHandler.Get)

	apiServer.ListenAndServe()
}

func bail(stage string, err error) {
	slog.Error("server bailing out", slog.String("stage", stage), "error", err)
	os.Exit(1)
}
