package lpapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetDomain(siteId string) (*LpDomains, error) {

	url := fmt.Sprintf("https://api.liveperson.net/api/account/%s/service/baseURI?version=1.0", siteId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))

	var result LpDomains
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	fmt.Printf("Response: %+v\n", result)

	return &result, nil

}
