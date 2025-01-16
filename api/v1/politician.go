package v1

type ID struct {
	Value  string `json:",chardata" xml:",chardata"`
	Status string `json:"status" xml:"status,attr"`
}

type MdbName struct {
	Value  string `json:",chardata" xml:",chardata"`
	Status string `json:"status" xml:"status,attr"`
}

type MdbWahlkreis struct {
	MdbWahlkreisNummer string `json:"mdbWahlkreisNummer" xml:"mdbWahlkreisNummer"`
	MdbWahlkreisName   string `json:"mdbWahlkreisName" xml:"mdbWahlkreisName"`
}

type Politician struct {
	DokumentInfo DokumentInfo  `json:"dokumentInfo" xml:"dokumentInfo"`
	MdbInfo      PoliticianBio `json:"mdbInfo" xml:"mdbInfo"`
	MdbMedien    MDBMedien     `json:"mdbMedien" xml:"mdbMedien"`
}

type PoliticianBio struct {
	ArticleID                            string        `json:"articleId" xml:"articleId"`
	SourceURL                            string        `json:"sourceURL" xml:"sourceURL"`
	ID                                   ID            `json:"mdbID" xml:"mdbID"`
	ExitDate                             string        `json:"mdbAustrittsdatum,omitempty" xml:"mdbAustrittsdatum"`
	LastName                             string        `json:"mdbZuname" xml:"mdbZuname"`
	FirstName                            string        `json:"mdbVorname" xml:"mdbVorname"`
	NobilityTitle                        string        `json:"mdbAdelstitel,omitempty" xml:"mdbAdelstitel"`
	AcademicTitle                        string        `json:"mdbAkademischerTitel,omitempty" xml:"mdbAkademischerTitel"`
	LocationSuffix                       string        `json:"mdbOrtszusatz,omitempty" xml:"mdbOrtszusatz"`
	DateOfBirth                          string        `json:"mdbGeburtsdatum" xml:"mdbGeburtsdatum"`
	ReligionOrDenomination               string        `json:"mdbReligionKonfession,omitempty" xml:"mdbReligionKonfession"`
	EducationOrProfessionalQualification string        `json:"mdbSchulOderBerufsabschluss,omitempty" xml:"mdbSchulOderBerufsabschluss"`
	HigherEducation                      string        `json:"mdbHochschulbildung,omitempty" xml:"mdbHochschulbildung"`
	Profession                           Profession    `json:"mdbBeruf" xml:"mdbBeruf"`
	Gender                               string        `json:"mdbGeschlecht" xml:"mdbGeschlecht"`
	MaritalStatus                        string        `json:"mdbFamilienstand,omitempty" xml:"mdbFamilienstand"`
	NumberKids                           string        `json:"mdbAnzahlKinder,omitempty" xml:"mdbAnzahlKinder"`
	Faction                              string        `json:"mdbFraktion" xml:"mdbFraktion"`
	Party                                string        `json:"mdbPartei" xml:"mdbPartei"`
	State                                string        `json:"mdbLand" xml:"mdbLand"`
	Constituency                         Constituency  `json:"mdbWahlkreis" xml:"mdbWahlkreis"`
	Elected                              string        `json:"mdbGewaehlt" xml:"mdbGewaehlt"`
	BioURL                               string        `json:"mdbBioURL" xml:"mdbBioURL"`
	BiographicInfo                       string        `json:"mdbBiografischeInformationen,omitempty" xml:"mdbBiografischeInformationen"`
	Trivia                               string        `json:"mdbWissenswertes,omitempty" xml:"mdbWissenswertes"`
	Homepage                             string        `json:"mdbHomepageURL,omitempty" xml:"mdbHomepageURL"`
	OtherWebsite                         OtherWebsites `json:"mdbSonstigeWebsites,omitempty" xml:"mdbSonstigeWebsites"`
	Phone                                string        `json:"mdbTelefon,omitempty" xml:"mdbTelefon"`
	Memberships                          Memberships   `json:"mdbMitgliedschaften,omitempty" xml:"mdbMitgliedschaften"`
	MandatedPublishableInfo              string        `json:"mdbVeroeffentlichungspflichtigeAngaben,omitempty" xml:"mdbVeroeffentlichungspflichtigeAngaben"`
}

type Profession struct {
	Field string `json:"berufsfeld,omitempty" xml:"berufsfeld,attr"`
	Value string `json:",chardata" xml:",chardata"`
}

type Constituency struct {
	Number string `json:"mdbWahlkreisNummer" xml:"mdbWahlkreisNummer"`
	Name   string `json:"mdbWahlkreisName" xml:"mdbWahlkreisName"`
	Url    string `json:"mdbWahlkreisURL" xml:"mdbWahlkreisURL"`
}

type OtherWebsites struct {
	Website []Website `json:"mdbSonstigeWebsite,omitempty" xml:"mdbSonstigeWebsite"`
}

type Website struct {
	Title string `json:"mdbSonstigeWebsiteTitel" xml:"mdbSonstigeWebsiteTitel"`
	URL   string `json:"mdbSonstigeWebsiteURL" xml:"mdbSonstigeWebsiteURL"`
}

type Memberships struct {
	LeadCommittees             []Committee `json:"mdbObleuteGremien,omitempty" xml:"mdbObleuteGremien>mdbObleuteGremium"`
	RegularMemberCommittees    []Committee `json:"mdbOrdentlichesMitgliedGremien,omitempty" xml:"mdbOrdentlichesMitgliedGremien>mdbOrdentlichesMitgliedGremium"`
	SubstituteMemberCommittees []Committee `json:"mdbStellvertretendesMitgliedGremien,omitempty" xml:"mdbStellvertretendesMitgliedGremien>mdbStellvertretendesMitgliedGremium"`
	ViceChairOtherCommittees   []Committee `json:"mdbStellvVorsitzSonstigeGremien,omitempty" xml:"mdbStellvVorsitzSonstigeGremien>mdbStellvVorsitzSonstigesGremium"`
}

type Committee struct {
	Id   string `json:"id,omitempty" xml:"id,attr"`
	Name string `json:"gremiumName" xml:"gremiumName"`
	Url  string `json:"gremiumURL" xml:"gremiumURL"`
}

type MDBMedien struct {
	Foto        Foto   `json:"mdbFoto" xml:"mdbFoto"`
	SpeechesUrl string `json:"mdbRedenVorPlenumURL" xml:"mdbRedenVorPlenumURL"`
	SpeechesRSS string `json:"mdbRedenVorPlenumRSS" xml:"mdbRedenVorPlenumRSS"`
}

type Foto struct {
	URL       string `json:"mdbFotoURL" xml:"mdbFotoURL"`
	Copyright string `json:"mdbFotoCopyright" xml:"mdbFotoCopyright"`
	LargeUrl  string `json:"mdbFotoGrossURL" xml:"mdbFotoGrossURL"`
	AltText   string `json:"imageAltText" xml:"imageAltText"`
}
