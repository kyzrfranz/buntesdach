package v1

type CommitteeDetails struct {
	DocumentInfo       DocumentInfo      `xml:"dokumentInfo" json:"document_info"`
	Id                 string            `xml:"ausschussId" json:"committee_id"`
	CommitteeName      string            `xml:"ausschussName" json:"committee_name"`
	SourceURL          string            `xml:"ausschussSourceURL" json:"source_url"`
	ImageURL           string            `xml:"ausschussBildURL" json:"image_url"`
	LargeImageURL      string            `xml:"ausschussBildGrossURL" json:"large_image_url"`
	ImageCopyright     string            `xml:"ausschussBildCopyright" json:"image_copyright"`
	ImageAltText       string            `xml:"imageAltText" json:"image_alt_text"`
	Tasks              string            `xml:"ausschussAufgabe" json:"tasks"`
	Contact            string            `xml:"ausschussKontakt" json:"contact"`
	ChairpersonID      string            `xml:"ausschussVorsitzId" json:"chairperson_id"`
	DeputyChairpersons []string          `xml:"ausschussStellvertretendeVorsitz>ausschussStellvertretendeVorsitzId" json:"deputy_chairpersons"`
	Members            []PersonListEntry `xml:"ausschussMitglieder>mdb" json:"members"`
	NewsItems          []NewsItem        `xml:"news>item" json:"news_items"`
}

type NewsItem struct {
	Title           string `xml:"title" json:"title"`
	Description     string `xml:"description" json:"description"`
	PublicationDate string `xml:"publicationDate" json:"publication_date"`
	URL             string `xml:"url" json:"url"`
}
