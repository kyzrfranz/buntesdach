package v1

type CommitteeDetails struct {
	DocumentInfo       DocumentInfo `xml:"dokumentInfo" json:"document_info"`
	CommitteeID        string       `xml:"ausschussId" json:"committee_id"`
	CommitteeName      string       `xml:"ausschussName" json:"committee_name"`
	SourceURL          string       `xml:"ausschussSourceURL" json:"source_url"`
	ImageURL           string       `xml:"ausschussBildURL" json:"image_url"`
	LargeImageURL      string       `xml:"ausschussBildGrossURL" json:"large_image_url"`
	ImageCopyright     string       `xml:"ausschussBildCopyright" json:"image_copyright"`
	ImageAltText       string       `xml:"imageAltText" json:"image_alt_text"`
	Tasks              string       `xml:"ausschussAufgabe" json:"tasks"`
	Contact            string       `xml:"ausschussKontakt" json:"contact"`
	ChairpersonID      string       `xml:"ausschussVorsitzId" json:"chairperson_id"`
	DeputyChairpersons []string     `xml:"ausschussStellvertretendeVorsitz>ausschussStellvertretendeVorsitzId" json:"deputy_chairpersons"`
	Members            []Member     `xml:"ausschussMitglieder>mdb" json:"members"`
	NewsItems          []NewsItem   `xml:"news>item" json:"news_items"`
}

type DocumentInfo struct {
	DocumentURL    string `xml:"dokumentURL" json:"document_url"`
	DocumentStatus string `xml:"dokumentStand" json:"document_status"`
}

type Member struct {
	Fraction        string `xml:"fraktion,attr" json:"fraction"`
	ID              string `xml:"mdbID" json:"id"`
	Name            string `xml:"mdbName" json:"name"`
	BioURL          string `xml:"mdbBioURL" json:"bio_url"`
	InfoXMLURL      string `xml:"mdbInfoXMLURL" json:"info_xml_url"`
	Region          string `xml:"mdbLand" json:"region"`
	Role            string `xml:"role" json:"role"`
	LastChanged     string `xml:"lastChanged" json:"last_changed"`
	ChangedDateTime string `xml:"changedDateTime" json:"changed_date_time"`
	PhotoURL        string `xml:"mdbFotoURL" json:"photo_url"`
	LargePhotoURL   string `xml:"mdbFotoGrossURL" json:"large_photo_url"`
	AltText         string `xml:"imageAltText" json:"alt_text"`
}

type NewsItem struct {
	Title           string `xml:"title" json:"title"`
	Description     string `xml:"description" json:"description"`
	PublicationDate string `xml:"publicationDate" json:"publication_date"`
	URL             string `xml:"url" json:"url"`
}
