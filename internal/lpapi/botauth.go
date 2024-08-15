package lpapi

import (
	"encoding/json"
	"io"
	"net/http"
    "strconv"

	"github.com/charmbracelet/log"
	"github.com/wlai-lp/bo-botflow/util"
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

func GetBotAccessToken(lpd *LpDomains, bearer string) (token string, orgId string, e error) {

	log.Info("getting access token with ", "bearer", bearer)
	uri, _ := getBaseURI(lpd, "cbLeIntegrations")

	uri = "https://" + uri + "/sso/authenticate"
	log.Debug("cbLeIntegrations domain is", "url", uri)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, uri, nil)

	if err != nil {
		log.Error(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Authorization", "Bearer "+bearer)

	res, err := client.Do(req)

    log.Info(res.StatusCode)
    if res.StatusCode !=200 {        
        return "", "", util.LogAndReturnError("request status code is " + strconv.Itoa(res.StatusCode))
    }

	if err != nil {
		log.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debug(string(body))

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("Error:", err)
		return
	}

    if result.SuccessResult.ApiAccessToken == ""{
        // return "", "", util.LogAndReturnError("access token is blank")
        return
    }
	log.Info("access token is:", "token", result.SuccessResult.ApiAccessToken)

	return result.SuccessResult.ApiAccessToken, result.SuccessResult.ChatBotPlatformUser.OrgId, nil
	// return ""
}
