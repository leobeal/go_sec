package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"stocks/types"
)

func FetchCompanies(searchTerm string) (types.Results, error) {

	url := "https://efts.sec.gov/LATEST/search-index"

	// JSON body
	body := []byte(`{
        "keysTyped": "` + searchTerm + `"
    }`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := makeRequest(req)
	defer res.Body.Close()

	if err != nil {
		return types.Results{}, err
	}

	results := types.Results{}
	err = json.NewDecoder(res.Body).Decode(&results)

	if err != nil {
		panic(err)
	}

	return results, nil
}

func FetchCompanyInfo(cik string) (types.Submission, error) {
	url := "https://data.sec.gov/submissions/CIK000" + cik + ".json"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Leonardo Beal leonardobeal@gmail.com")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Host", "www.sec.gov")

	res, err := makeRequest(req)
	defer res.Body.Close()

	if err != nil {
		return types.Submission{}, err
	}

	submissions := types.Submission{}
	err = json.NewDecoder(res.Body).Decode(&submissions)

	if err != nil {
		panic(err)
	}

	return submissions, nil
}

func makeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)

	return res, err
}
