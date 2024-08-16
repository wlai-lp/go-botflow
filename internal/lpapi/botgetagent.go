// this returns the list of both within the group
// this is the same as ungroup and groupid call
// just query param diff
// bot-group=un_assigned
// bot-group=whatevergroupid

package lpapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
)

// Define the structs
type BotAgentResultConfiguration struct {
    LpUserRole         string `json:"lpUserRole"`
    EnableAccessibility string `json:"enableAccessibility"`
}

type BotAgentResultAgent struct {
    CreateTime                  int64        `json:"createTime"`
    UpdateTime                  int64        `json:"updateTime"`
    AuthType                    string       `json:"authType"`
    Success                     bool         `json:"success"`
    PipelineConfig              interface{}  `json:"pipelineConfig"`
    LpAccountId                 string       `json:"lpAccountId"`
    LpAccountUser               string       `json:"lpAccountUser"`
    LpUserId                    string       `json:"lpUserId"`
    Type                        string       `json:"type"`
    BotId                       string       `json:"botId"`
    ServerName                  string       `json:"serverName"`
    Configurations              BotAgentResultConfiguration `json:"configurations"`
    LastManualOperation         string       `json:"lastManualOperation"`
    ManualOperationTs           int64        `json:"manualOperationTs"`
    ManualOperationPerformedById string      `json:"manualOperationPerformedById"`
    ManualOperationPerformedByName string    `json:"manualOperationPerformedByName"`
    DeploymentEnvironment       string       `json:"deploymentEnvironment"`
}

type BotAgentResult struct {
    ActivationId string  `json:"activationId"`
    ChatBotId    string  `json:"chatBotId"`
    CompanyName  string  `json:"companyName"`
    CreateTime   int64   `json:"createTime"`
    UpdateTime   int64   `json:"updateTime"`
    Status       string  `json:"status"`
    Agents       []BotAgentResultAgent `json:"agents"`
}

// return 1 agent is enough
func GetBotAgentByBotId(lpd *LpDomains, token string, orgId string, botId string) string {

	log.Info("GetBotGroups", "token", token)
	uri, _ := getBaseURI(lpd, "cbBotPlatform")

	// "https://va.bc-platform.liveperson.net/bot-platform-manager-0.1/bot-groups/bots?sort-by=botName%3Aasc&size=50&bot-group=un_assigned"
	

	// https://va.bc-platform.liveperson.net/bot-platform-manager-0.1/auth/liveperson/app?chatBotId=02bcaffc-fb22-4282-9208-bbcf9d0f7d0e
	// https://va.bc-platform.liveperson.net/bot-platform-manager-0.1/auth/liveperson/app?chatBotId=e9d8f6a8-5d7e-496b-98b2-1b37a8c0a080"
	url := "https://" + uri + "/auth/liveperson/app?chatBotId=" + botId
	log.Info("get", "url", url)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("OrganizationId", orgId)

	res, err := client.Do(req)

	if res.StatusCode == 404 {
		log.Warn("No bot agent found")
		return ""
	}

	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	log.Info(string(body))

	var result BotAgentResult
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("Error:", err)
		return ""
	}
	log.Info("BotAgent by botId agent", "count", len(result.Agents))
	return result.Agents[0].LpAccountUser
}
