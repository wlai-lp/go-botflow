package lpapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SuccessResult struct {
	ChatBotPlatformUser   ChatBotPlatformUser `json:"chatBotPlatformUser"`
	ApiAccessToken        string              `json:"apiAccessToken"`
	ApiAccessPermToken    string              `json:"apiAccessPermToken"`
	Config                Config              `json:"config"`
	SessionOrganizationId string              `json:"sessionOrganizationId"`
	LeAccountId           string              `json:"leAccountId"`
	CbRegion              string              `json:"cbRegion"`
	EnabledFeatures       []string            `json:"enabledFeatures"`
	SiteSettings          []SiteSetting       `json:"siteSettings"`
	LeUserId              string              `json:"leUserId"`
	Privileges            []int               `json:"privileges"`
	IsElevatedLpa         bool                `json:"isElevatedLpa"`
}

type ChatBotPlatformUser struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	FirstName        string `json:"firstname"`
	LastName         string `json:"lastname"`
	UserLoginType    string `json:"userLoginType"`
	UserId           string `json:"userId"`
	Role             string `json:"role"`
	OrgId            string `json:"orgId"`
	Status           string `json:"status"`
	CreationTime     string `json:"creationTime"`
	ModificationTime string `json:"modificationTime"`
	Cb2Enabled       bool   `json:"cb2Enabled"`
}

type Config struct {
	TrainMinSizeSamples string `json:"train.min_size.samples"`
	TrainMinSizeIntents string `json:"train.min_size.intents"`
}

type SiteSetting struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type Response struct {
	Success       bool          `json:"success"`
	SuccessResult SuccessResult `json:"successResult"`
	Message       string        `json:"message"`
}

func GetBotAccessToken(lpd *LpDomains, bearer string) (token string) {

    // fmt.Printf("\ngetting access token with bearer %v\n", bearer)
	// url := "https://va.bc-sso.liveperson.net/le-auth/sso/authenticate"
	// method := "GET"

	// client := &http.Client{}
	// req, err := http.NewRequest(method, url, nil)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// req.Header.Add("Accept", "application/json, text/plain, */*")
	// // req.Header.Add("Authorization", "Bearer 9e23d9cc84188a48380010a4bedbb298071de032afd8051e31ad93439f895d35")
	// req.Header.Add("Authorization", "Bearer " + bearer)

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))

	fmt.Printf("\ngetting access token with bearer %v\n", bearer)
	uri, _ := getBaseURI(lpd, "cbLeIntegrations")

	uri = "https://" + uri + "/sso/authenticate"
	fmt.Printf("uri is %v \n", uri)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, uri, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Authorization", "Bearer " + bearer)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
	    fmt.Println("\nError:", err)
		return ""
	}
	fmt.Printf("\naccess token is: %v\n", result.SuccessResult.ApiAccessToken)

	return result.SuccessResult.ApiAccessToken
    // return ""
}
