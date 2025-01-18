package config

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagNameDataURL      = "data-url"
	flagNameCommitteeURL = "committee-url"
	flagServerPort       = "server-port"

	defaultDataURL      = "https://www.bundestag.de/xml/v2/mdb/index.xml"
	defaultCommitteeURL = "https://www.bundestag.de/xml/v2/ausschuesse/index.xml"
	defaultServerPort   = 8080
)

type Config struct {
	DataURL      *url.URL
	CommitteeURL *url.URL
	ServerPort   int
}

// Seed is seeding a `Config` from command line arguments and environments variables.
// Command line take precedence environments variables.
func Seed() (*Config, error) {
	// define defaults
	viper.SetDefault(flagNameDataURL, defaultDataURL)
	viper.SetDefault(flagNameCommitteeURL, defaultCommitteeURL)
	viper.SetDefault(flagServerPort, defaultServerPort)

	// setup cmd line flags
	pflag.String(flagNameDataURL, defaultDataURL, "Data URL")
	pflag.String(flagNameCommitteeURL, defaultCommitteeURL, "Committee URL")
	pflag.Int(flagServerPort, defaultServerPort, "API Server port")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return nil, fmt.Errorf("failed to bind flags: %w", err)
	}

	// setup env vars
	viper.SetEnvPrefix("BUNTESDACH")
	r := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(r) // NOTE: We want BUNTESDACH_SERVER_PORT not BUNTESDACH_SERVER-PORT
	viper.BindEnv(flagNameDataURL)
	viper.BindEnv(flagNameCommitteeURL)
	viper.BindEnv(flagServerPort)

	// build config
	dataURL, err := mustGetUrl(viper.GetString(flagNameDataURL))
	if err != nil {
		return nil, fmt.Errorf("data URL: %w", err)
	}

	committeeURL, err := mustGetUrl(viper.GetString(flagNameDataURL))
	if err != nil {
		return nil, fmt.Errorf("committee URL: %w", err)
	}

	return &Config{
		DataURL:      dataURL,
		CommitteeURL: committeeURL,
		ServerPort:   viper.GetInt(flagServerPort),
	}, nil
}

func mustGetUrl(s string) (*url.URL, error) {
	parsedUrl, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("parse data URL: %w", err)
	}

	return parsedUrl, err
}
