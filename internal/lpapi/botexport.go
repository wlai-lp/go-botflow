// this returns the list of both within the group
// this is the same as ungroup and groupid call
// just query param diff
// bot-group=un_assigned
// bot-group=whatevergroupid

package lpapi

import (
	// "encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	// tea "github.com/charmbracelet/bubbletea"
	"github.com/wlai-lp/bo-botflow/util"
)

func TestMethod() string{
	// log.Info("logging")
	// return func() tea.Msg {
		return "executed Test method"
	// }
}

func GetBotByBotIdTeaCmd(botId string) string {
	err:=GetBotByBotId(&APP_PARAM.DOMAINS, APP_PARAM.ACCESS_TOKEN, APP_PARAM.ORG_ID, botId)	
	if err != nil {
		return "ERROR: unable to get bot"
	}
	return "Bot saved to file"
}

func GetBotByBotId(lpd *LpDomains, token string, orgId string, botId string) error {

	
	log.Info("GetBotByBotIdFull", "token", token, "orgId", orgId)
	uri, _ := getBaseURI(lpd, "cbBotPlatform")

	// https://va.bc-platform.liveperson.net/bot-platform-manager-0.1/tools/export/9cef068d-2f16-4437-8a22-ff24a84f8bf7
	// https://va.bc-platform.liveperson.net/bot-platform-manager-0.1/tools/export/9cef068d-2f16-4437-8a22-ff24a84f8bf7	   
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

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
	// log.Info(string(body))

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

	// var result BotGroupResponse
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	log.Error("Error:", err)
	// 	return nil
	// }

	log.Info("bot saved to local file")
	return nil
}
