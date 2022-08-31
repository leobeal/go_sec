package types

import "time"

type Results struct {
	Took int       `json:"took"`
	Hits HitParent `json:"hits"`
}

type HitParent struct {
	Hits []HitChild `json:"hits"`
}

type HitChild struct {
	Id     string `json:"_id"`
	Index  string `json:"_index"`
	Source Source `json:"_source"`
}

type Source struct {
	Entity string `json:"entity"`
	Ticker string `json:"tickers"`
	Rank   int    `json:"rank"`
}

type Submission struct {
	Cik                               string   `json:"cik"`
	EntityType                        string   `json:"entityType"`
	Sic                               string   `json:"sic"`
	SicDescription                    string   `json:"sicDescription"`
	InsiderTransactionForOwnerExists  int      `json:"insiderTransactionForOwnerExists"`
	InsiderTransactionForIssuerExists int      `json:"insiderTransactionForIssuerExists"`
	Name                              string   `json:"name"`
	Ticker                            []string `json:"tickers"`
	Exchanges                         []string `json:"exchanges"`
	Ein                               string   `json:"ein"`
	Description                       string   `json:"description"`
	Website                           string   `json:"website"`
	InvestorWebsite                   string   `json:"investorWebsite"`
	Category                          string   `json:"category"`
	FiscalYearEnd                     string   `json:"fiscalYearEnd"`
	StateOfIncorporation              string   `json:"stateOfIncorporation"`
	StateOfIncorporationDescription   string   `json:"stateOfIncorporationDescription"`
	Addresses                         struct {
		Mailing struct {
			Street1                   string      `json:"street1"`
			Street2                   interface{} `json:"street2"`
			City                      string      `json:"city"`
			StateOrCountry            string      `json:"stateOrCountry"`
			ZipCode                   string      `json:"zipCode"`
			StateOrCountryDescription string      `json:"stateOrCountryDescription"`
		} `json:"mailing"`
		Business struct {
			Street1                   string      `json:"street1"`
			Street2                   interface{} `json:"street2"`
			City                      string      `json:"city"`
			StateOrCountry            string      `json:"stateOrCountry"`
			ZipCode                   string      `json:"zipCode"`
			StateOrCountryDescription string      `json:"stateOrCountryDescription"`
		} `json:"business"`
	} `json:"addresses"`
	Phone       string `json:"phone"`
	Flags       string `json:"flags"`
	FormerNames []struct {
		Name string    `json:"name"`
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	} `json:"formerNames"`
	Filings struct {
		Recent Filling       `json:"recent"`
		Files  []interface{} `json:"files"`
	} `json:"filings"`
}

type Filling struct {
	AccessionNumber       []string    `json:"accessionNumber"`
	FilingDate            []string    `json:"filingDate"`
	ReportDate            []string    `json:"reportDate"`
	AcceptanceDateTime    []time.Time `json:"acceptanceDateTime"`
	Act                   []string    `json:"act"`
	Form                  []string    `json:"form"`
	FileNumber            []string    `json:"fileNumber"`
	FilmNumber            []string    `json:"filmNumber"`
	Items                 []string    `json:"items"`
	Size                  []int       `json:"size"`
	IsXBRL                []int       `json:"isXBRL"`
	IsInlineXBRL          []int       `json:"isInlineXBRL"`
	PrimaryDocument       []string    `json:"primaryDocument"`
	PrimaryDocDescription []string    `json:"primaryDocDescription"`
}
