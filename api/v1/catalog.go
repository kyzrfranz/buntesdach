package v1

type PersonCatalog struct {
	DokumentInfo  DokumentInfo      `json:"dokumentInfo" xml:"dokumentInfo"`
	DeleteRestore DeleteRestore     `json:"deleteRestore" xml:"deleteRestore"`
	Persons       []PersonListEntry `json:"mdbs" xml:"mdbs>mdb"`
}

type DeleteRestore struct {
	DeleteFlag string `json:"deleteFlag" xml:"deleteFlag"`
	DeleteDate string `json:"deleteDate" xml:"deleteDate"`
}

type PersonListEntry struct {
	Faction              string       `json:"fraktion" xml:"fraktion,attr"`
	Id                   ID           `json:"mdbID" xml:"mdbID"`
	Name                 MdbName      `json:"mdbName" xml:"mdbName"`
	BioURL               string       `json:"mdbBioURL" xml:"mdbBioURL"`
	InfoXMLURL           string       `json:"mdbInfoXMLURL" xml:"mdbInfoXMLURL"`
	InfoXMLURLMitmischen string       `json:"mdbInfoXMLURLMitmischen,omitempty" xml:"mdbInfoXMLURLMitmischen"`
	State                string       `json:"mdbLand" xml:"mdbLand"`
	Constituancy         MdbWahlkreis `json:"mdbWahlkreis" xml:"mdbWahlkreis"`
	Elected              string       `json:"mdbGewaehlt" xml:"mdbGewaehlt"`
	FotoURL              string       `json:"mdbFotoURL" xml:"mdbFotoURL"`
	FotoGrossURL         string       `json:"mdbFotoGrossURL" xml:"mdbFotoGrossURL"`
	FotoLastChanged      string       `json:"mdbFotoLastChanged" xml:"mdbFotoLastChanged"`
	FotoChangedDateTime  string       `json:"mdbFotoChangedDateTime" xml:"mdbFotoChangedDateTime"`
	ImageAltText         string       `json:"imageAltText" xml:"imageAltText"`
	LastChanged          string       `json:"lastChanged" xml:"lastChanged"`
	ChangedDateTime      string       `json:"changedDateTime" xml:"changedDateTime"`
}
