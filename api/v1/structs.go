package v1

type Politician struct {
	Img           string `json:"img"`
	Name          string `json:"name"`
	Direct        int    `json:"direct"`
	Gender        string `json:"geschlecht"`
	Id            string `json:"id"`
	Href          string `json:"href"`
	AgeGroup      string `json:"ageGroup"`
	LastnameFirst string `json:"lastnamefirst"`
	FederalState  string `json:"federalState"`
	Party         string `json:"party"`
}

type PoliticianCatalog map[string]Politician
