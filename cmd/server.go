package main

import (
	"log/slog"
	"os"

	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/kyzrfranz/buntesdach/internal/config"
	"github.com/kyzrfranz/buntesdach/internal/data"
	"github.com/kyzrfranz/buntesdach/internal/http"
	"github.com/kyzrfranz/buntesdach/internal/rest"
	"github.com/kyzrfranz/buntesdach/internal/upstream"
	"github.com/kyzrfranz/buntesdach/pkg/resources"
)

func main() {
	config, err := config.Seed()
	if err != nil {
		bail("Seeding config", err)
	}

	slog.Info("Config",
		"dataURL", config.DataURL.String(),
		"committeeURL", config.CommitteeURL.String(),
		"serverPort", config.ServerPort,
	)

	politicianReader := data.NewCatalogReader[v1.PersonCatalog](&upstream.XMLFetcher{Url: config.DataURL})
	committeeReader := data.NewCatalogReader[v1.CommitteeCatalog](&upstream.XMLFetcher{Url: config.CommitteeURL})

	apiServer := http.NewApiServer(config.ServerPort)

	apiServer.Use(http.MiddlewareRecovery)
	apiServer.Use(http.MiddlewareCORS)

	politicianCatalogHandler := rest.NewHandler(resources.NewCatalogueRepo(&politicianReader))
	politicianDetailHandler := rest.NewHandler(resources.NewDetailRepo[v1.Politician](&politicianReader))
	committeeCatalogueHandler := rest.NewHandler(resources.NewCatalogueRepo(&committeeReader))
	committeeDetailHandler := rest.NewHandler(resources.NewDetailRepo[v1.CommitteeDetails](&committeeReader))

	apiServer.AddHandler("/politicians", politicianCatalogHandler.List)
	apiServer.AddHandler("/politicians/{id}", politicianCatalogHandler.Get)
	apiServer.AddHandler("/politicians/{id}/bio", politicianDetailHandler.Get)
	apiServer.AddHandler("/committees", committeeCatalogueHandler.List)
	apiServer.AddHandler("/committees/{id}", committeeCatalogueHandler.Get)
	apiServer.AddHandler("/committees/{id}/detail", committeeDetailHandler.Get)

	apiServer.ListenAndServe()
}

func bail(stage string, err error) {
	slog.Error("server bailing out", slog.String("stage", stage), "error", err)
	os.Exit(1)
}
