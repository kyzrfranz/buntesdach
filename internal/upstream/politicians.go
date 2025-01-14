package upstream

import (
	"encoding/json"
	"fmt"
	myhttp "github.com/kyzrfranz/buntesdach/internal/http"
	"net/url"
)

type Politicians struct {
	Url *url.URL
}

func (p *Politicians) Fetch() ([]byte, error) {
	data, err := myhttp.FetchUrl(p.Url)
	if err != nil {
		return nil, err
	}

	// Optionally, validate that the response is valid JSON
	var jsonData interface{}
	if err = json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	return data, nil
}
