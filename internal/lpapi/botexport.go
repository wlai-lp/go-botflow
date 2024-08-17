package lpapi

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/wlai-lp/bo-botflow/util"
)

func ExportBotByBotIdTeaCmd(botId string) string {
	err := ExportBotByBotId(&APP_PARAM.DOMAINS, APP_PARAM.ACCESS_TOKEN, APP_PARAM.ORG_ID, botId)
	if err != nil {
		return "ERROR: unable to get bot"
	}
	return "Bot saved to file"
}

func ExportBotByBotId(lpd *LpDomains, token string, orgId string, botId string) error {

	log.Info("GetBotByBotIdFull", "token", token, "orgId", orgId)
	uri, _ := getBaseURI(lpd, "cbBotPlatform")

	url := "https://" + uri + "/tools/export/" + botId
	log.Info("GetBotByBotIdFull", "url", url)
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

	// Create the file
	file, err := os.Create(fmt.Sprintf("%v.json", botId))
	if err != nil {
		util.LogAndReturnError("unable to create file")
	}
	defer file.Close()

	// Write the response body to the file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		util.LogAndReturnError("unable to save file")
	}


	log.Info("bot saved to local file")
	return nil
}
