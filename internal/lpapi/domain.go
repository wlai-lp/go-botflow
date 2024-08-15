package lpapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	 "github.com/charmbracelet/log"
)

type LpDomains struct {
	BaseURIs []struct {
		Service string `json:"service"`
		Account string `json:"account"`
		BaseURI string `json:"baseURI"`
	} `json:"baseURIs"`
}

func Hello() string {
	return "hello"
}

func getBaseURI(domains *LpDomains, service string) (string, bool) {
	for _, uri := range domains.BaseURIs {
		if uri.Service == service {
			return uri.BaseURI, true
		}
	}
	return "", false
}



func GetDomain(siteId string) (*LpDomains, error) {

	url := fmt.Sprintf("https://api.liveperson.net/api/account/%s/service/baseURI?version=1.0", siteId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debug(string(body))

	var result LpDomains
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("Error:", err)
		return nil, err
	}

	return &result, nil

}
