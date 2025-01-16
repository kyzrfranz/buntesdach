package v1

type DokumentInfo struct {
	DokumentURL   string `json:"dokumentURL,omitempty" xml:"dokumentURL"`
	DokumentStand string `json:"dokumentStand" xml:"dokumentStand"`
}
