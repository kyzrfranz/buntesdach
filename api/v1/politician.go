package v1

type ID struct {
	Value  string `json:"value" xml:",chardata"`
	Status string `json:"status" xml:"status,attr"`
}

type MdbName struct {
	Value  string `json:"value" xml:",chardata"`
	Status string `json:"status" xml:"status,attr"`
}

type Politician struct {
	DocumentInfo DocumentInfo  `json:"documentInfo" xml:"dokumentInfo"`
	Bio          PoliticianBio `json:"bio" xml:"mdbInfo"`
	Media        Media         `json:"media" xml:"mdbMedien"`
}

type PoliticianBio struct {
	Id                                   ID            `json:"id" xml:"mdbID"`
	ArticleID                            string        `json:"articleId" xml:"articleId"`
	SourceURL                            string        `json:"sourceUrl" xml:"sourceURL"`
	ExitDate                             string        `json:"exitDate,omitempty" xml:"mdbAustrittsdatum"`
	LastName                             string        `json:"lastName" xml:"mdbZuname"`
	FirstName                            string        `json:"firstName" xml:"mdbVorname"`
	NobilityTitle                        string        `json:"nobilityTitle,omitempty" xml:"mdbAdelstitel"`
	AcademicTitle                        string        `json:"academicTitle,omitempty" xml:"mdbAkademischerTitel"`
	LocationSuffix                       string        `json:"locationSuffix,omitempty" xml:"mdbOrtszusatz"`
	DateOfBirth                          string        `json:"dateOfBirth" xml:"mdbGeburtsdatum"`
	ReligionOrDenomination               string        `json:"religionOrDenomination,omitempty" xml:"mdbReligionKonfession"`
	EducationOrProfessionalQualification string        `json:"educationOrProfessionalQualification,omitempty" xml:"mdbSchulOderBerufsabschluss"`
	HigherEducation                      string        `json:"higherEducation,omitempty" xml:"mdbHochschulbildung"`
	Profession                           Profession    `json:"profession" xml:"mdbBeruf"`
	Gender                               string        `json:"gender" xml:"mdbGeschlecht"`
	MaritalStatus                        string        `json:"maritalStatus,omitempty" xml:"mdbFamilienstand"`
	NumberKids                           string        `json:"numberOfKids,omitempty" xml:"mdbAnzahlKinder"`
	Faction                              string        `json:"faction" xml:"mdbFraktion"`
	Party                                string        `json:"party" xml:"mdbPartei"`
	State                                string        `json:"state" xml:"mdbLand"`
	Constituency                         Constituency  `json:"constituency" xml:"mdbWahlkreis"`
	Elected                              string        `json:"elected" xml:"mdbGewaehlt"`
	BioURL                               string        `json:"bioUrl" xml:"mdbBioURL"`
	BiographicInfo                       string        `json:"biographicInfo,omitempty" xml:"mdbBiografischeInformationen"`
	Trivia                               string        `json:"trivia,omitempty" xml:"mdbWissenswertes"`
	Homepage                             string        `json:"homepage,omitempty" xml:"mdbHomepageURL"`
	OtherWebsite                         OtherWebsites `json:"otherWebsites,omitempty" xml:"mdbSonstigeWebsites"`
	Phone                                string        `json:"phone,omitempty" xml:"mdbTelefon"`
	Memberships                          Memberships   `json:"memberships,omitempty" xml:"mdbMitgliedschaften"`
	MandatedPublishableInfo              string        `json:"mandatedPublishableInfo,omitempty" xml:"mdbVeroeffentlichungspflichtigeAngaben"`
}

type Profession struct {
	Field string `json:"field,omitempty" xml:"berufsfeld,attr"`
	Value string `json:"value" xml:",chardata"`
}

type Constituency struct {
	Number string `json:"number,omitempty" xml:"mdbWahlkreisNummer"`
	Name   string `json:"name,omitempty" xml:"mdbWahlkreisName"`
	Url    string `json:"url,omitempty" xml:"mdbWahlkreisURL"`
}

type OtherWebsites struct {
	Website []Website `json:"websites,omitempty" xml:"mdbSonstigeWebsite"`
}

type Website struct {
	Title string `json:"title" xml:"mdbSonstigeWebsiteTitel"`
	URL   string `json:"url" xml:"mdbSonstigeWebsiteURL"`
}

type Memberships struct {
	LeadCommittees             []Committee `json:"leadCommittees,omitempty" xml:"mdbObleuteGremien>mdbObleuteGremium"`
	RegularMemberCommittees    []Committee `json:"regularMemberCommittees,omitempty" xml:"mdbOrdentlichesMitgliedGremien>mdbOrdentlichesMitgliedGremium"`
	SubstituteMemberCommittees []Committee `json:"substituteMemberCommittees,omitempty" xml:"mdbStellvertretendesMitgliedGremien>mdbStellvertretendesMitgliedGremium"`
	ViceChairOtherCommittees   []Committee `json:"viceChairOtherCommittees,omitempty" xml:"mdbStellvVorsitzSonstigeGremien>mdbStellvVorsitzSonstigesGremium"`
}

type Committee struct {
	Id   string `json:"id,omitempty" xml:"id,attr"`
	Name string `json:"name" xml:"gremiumName"`
	Url  string `json:"url" xml:"gremiumURL"`
}

type Media struct {
	Foto        Foto   `json:"photo" xml:"mdbFoto"`
	SpeechesUrl string `json:"speechesUrl" xml:"mdbRedenVorPlenumURL"`
	SpeechesRSS string `json:"speechesRSS" xml:"mdbRedenVorPlenumRSS"`
}

type Foto struct {
	URL       string `json:"url" xml:"mdbFotoURL"`
	Copyright string `json:"copyright" xml:"mdbFotoCopyright"`
	LargeUrl  string `json:"largeUrl" xml:"mdbFotoGrossURL"`
	AltText   string `json:"altText" xml:"imageAltText"`
}
