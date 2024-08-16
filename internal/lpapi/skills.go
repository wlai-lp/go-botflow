// api to get skills
package lpapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
)

// Define the struct
type Skill struct {
    Deleted bool   `json:"deleted"`
    Name    string `json:"name"`
    ID      int64  `json:"id"`
}

func AAA(){
	
}

func GetSkills(lpd *LpDomains, siteId string, bearer string) []Skill {

	log.Info("Get Skills", "bearer", bearer)
	uri, _ := getBaseURI(lpd, "accountConfigReadWrite")

	// "https://va.bc-platform.liveperson.net/bot-platform-manager-0.1/bot-groups/bots?sort-by=botName%3Aasc&size=50&bot-group=un_assigned"
	
	// url := "https://" + uri + "/bot-groups/bots?sort-by=botName%3Aasc&size=50&bot-group=un_assigned"
	// url := "https://" + uri + "/bot-groups/bots?sort-by=botName%3Aasc&size=50&bot-group=" + groupid
	// "https://va.ac.liveperson.net/api/account/90412079/configuration/le-users/skills?v=2.0"
	url := "https://" + uri + "/api/account/" + siteId + "/configuration/le-users/skills?v=2.0"
	log.Info("get", "url", url)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authorization", "Bearer " + bearer)

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

	var result []Skill
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("Error:", err)
		return nil
	}
	log.Info("get skills count:", "result", len(result))
	return result
}
