package lpapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
)

// "encoding/json"
// "io"
// "net/http"
// "strconv"

// "github.com/charmbracelet/log"
// "github.com/wlai-lp/bo-botflow/util"

// Define the Data struct
type BotGroupsData struct {
	BotGroupID           string      `json:"botGroupId"`
	BotGroupName         string      `json:"botGroupName"`
	TransferMessage      interface{} `json:"transferMessage"`
	Channel              string      `json:"channel"`
	OrganizationID       interface{} `json:"organizationId"`
	CollaborationEnabled bool        `json:"collaborationEnabled"`
	NumberOfBots         int         `json:"numberOfBots"`
	CreatedAt            int         `json:"createdAt"`
	UpdatedAt            int         `json:"updatedAt"`
	CreatedBy            interface{} `json:"createdBy"`
	CreatedByName        interface{} `json:"createdByName"`
	UpdatedBy            interface{} `json:"updatedBy"`
	UpdatedByName        interface{} `json:"updatedByName"`
}

// Define the AutoGenerated struct
type BotGroupResult struct {
	Success       bool `json:"success"`
	SuccessResult struct {
		BotGroupsData []BotGroupsData `json:"data"`
	} `json:"successResult"`
}

func GetBotGroups(lpd *LpDomains, token string, orgId string) []BotGroupsData {

	log.Info("GetBotGroups", "token", token)
	uri, _ := getBaseURI(lpd, "cbBotPlatform")

	url := "https://" + uri + "/bot-groups?expand-all=true"
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
	log.Debug(string(body))

	var result BotGroupResult
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("Error:", err)
		return nil
	}
	log.Info("access token is:", "result", result.Success)
	return result.SuccessResult.BotGroupsData
}
