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

	politicianReader := data.NewCatalogReader[v1.PersonCatalog, v1.PersonListEntry](&upstream.XMLFetcher{Url: config.DataURL})
	committeeReader := data.NewCatalogReader[v1.CommitteeCatalog, v1.CommitteeListEntry](&upstream.XMLFetcher{Url: config.CommitteeURL})

	apiServer := http.NewApiServer(config.ServerPort)

	apiServer.Use(http.MiddlewareRecovery)
	apiServer.Use(http.MiddlewareCORS)

	politicianCatalogHandler := rest.NewHandler[v1.PersonListEntry](resources.NewCatalogueRepo[v1.PersonListEntry](&politicianReader))
	politicianDetailHandler := rest.NewHandler[v1.Politician](resources.NewDetailRepo[v1.Politician](&politicianReader))
	committeeCatalogueHandler := rest.NewHandler[v1.CommitteeListEntry](resources.NewCatalogueRepo[v1.CommitteeListEntry](&committeeReader))
	committeeDetailHandler := rest.NewHandler[v1.CommitteeDetails](resources.NewDetailRepo[v1.CommitteeDetails](&committeeReader))

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
