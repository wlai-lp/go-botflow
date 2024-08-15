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
type BotGroupPageContext struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalSize int `json:"totalSize"`
}

type GroupBot struct {
	BotID                   string  `json:"botId"`
	BotName                 string  `json:"botName"`
	BotDescription          string  `json:"botDescription"`
	BotType                 string  `json:"botType"`
	Channel                 string  `json:"channel"`
	BotLanguage             string  `json:"botLanguage"`
	AgentAnnotationsEnabled bool    `json:"agentAnnotationsEnabled"`
	DebuggingEnabled        bool    `json:"debuggingEnabled"`
	BotVersion              string  `json:"botVersion"`
	EntityDataSourceID      string  `json:"entityDataSourceId"`
	PublicBot               bool    `json:"publicBot"`
	OrganizationID          string  `json:"organizationId"`
	BotGroupID              string  `json:"botGroupId"`
	ChatBotPlatformUserID   string  `json:"chatBotPlatformUserId"`
	CreatedAt               int64   `json:"createdAt"`
	UpdatedAt               int64   `json:"updatedAt"`
	CreatedBy               *string `json:"createdBy"`
	CreatedByName           *string `json:"createdByName"`
	UpdatedBy               string  `json:"updatedBy"`
	UpdatedByName           string  `json:"updatedByName"`
	NumberOfDialogs         int     `json:"numberOfDialogs"`
	NumberOfInteractions    int     `json:"numberOfInteractions"`
	NumberOfIntegrations    int     `json:"numberOfIntegrations"`
	NumberOfActiveAgents    int     `json:"numberOfActiveAgents"`
	NumberOfInactiveAgents  int     `json:"numberOfInactiveAgents"`
	NumberOfDomains         int     `json:"numberOfDomains"`
	NumberOfIntents         int     `json:"numberOfIntents"`
	HasDisambiguation       bool    `json:"hasDisambiguation"`
	HasAutoescalation       bool    `json:"hasAutoescalation"`
	SmallTalkEnabled        bool    `json:"smallTalkEnabled"`
}

type BotGroupSuccessResult struct {
	PageContext BotGroupPageContext `json:"pageContext"`
	Data        []GroupBot          `json:"data"`
}

type BotGroupResponse struct {
	Success       bool                  `json:"success"`
	SuccessResult BotGroupSuccessResult `json:"successResult"`
}

func GetBotsByGroupId(lpd *LpDomains, token string, orgId string, groupid string) []GroupBot {

	log.Info("GetBotGroups", "token", token)
	uri, _ := getBaseURI(lpd, "cbBotPlatform")

	// "https://va.bc-platform.liveperson.net/bot-platform-manager-0.1/bot-groups/bots?sort-by=botName%3Aasc&size=50&bot-group=un_assigned"
	var postfix string
	if groupid == "un_assigned" {
		// postfix := "bot-group=un_assigned"
		// postfix := 
		postfix = "bot-group=un_assigned"
	} else {
		// postfix := "bot-group=un_assigned"
		postfix = "bot-group-id=" + groupid
	}

	// url := "https://" + uri + "/bot-groups/bots?sort-by=botName%3Aasc&size=50&bot-group=un_assigned"
	// url := "https://" + uri + "/bot-groups/bots?sort-by=botName%3Aasc&size=50&bot-group=" + groupid
	url := "https://" + uri + "/bot-groups/bots?sort-by=botName%3Aasc&size=50&" + postfix
	log.Info("get", "url", url)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("OrganizationId", orgId)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	log.Info(string(body))

	var result BotGroupResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("Error:", err)
		return nil
	}
	log.Info("access token is:", "result", result.Success)
	return result.SuccessResult.Data
}
