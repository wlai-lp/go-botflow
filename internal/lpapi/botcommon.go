package lpapi

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type Bot struct {
	ID     string
	Name   string
	Group  string
	Agents string
	Skills string
}

const UNASSIGNED = "un_assigned"

var groupNameByIdMap = make(map[string]string)
var botAgentsMap = make(map[string]string)
var agentSkillsMap = make(map[string]string)
var skillIdToNameMap = make(map[int64]string)

func GetListOfBots(siteId string, bearer string) ([]Bot, error) {
	start := time.Now()
	// check required fields
	if bearer == "" || siteId == "" {
		log.Error("missing site and/or bearer token value")
		return nil, errors.New("site/bearer token value is empty")
	}

	// get domain by siteid
	lpd, err := GetDomain(fmt.Sprint(viper.Get("LP_SITE")))
	if err != nil {
		return nil, err
	}

	// get bot access token
	token, orgId, err := GetBotAccessToken(lpd, bearer)
	if err != nil {
		return nil, err
	}
	log.Info(fmt.Sprintf("token is %v and org is %v", token, orgId))

	// get bot group list to get group id
	groups := GetBotGroups(lpd, token, orgId)
	log.Info("total of", "groups", len(groups))

	// cache groups id to name
	// groupNameById := make(map[string]string)
	for _, g := range groups {
		groupNameByIdMap[g.BotGroupID] = g.BotGroupName
	}

	// get bots by group id
	allBots := GetBotsByGroupId(lpd, token, orgId, UNASSIGNED)
	log.Info("total of ungrouped", "ungrouped", len(allBots))

	// loop each botgroup id to get bots in the group, add it to allbots
	// todo: make this call concurrent
	var wg sync.WaitGroup
	// results := make(chan []GroupBot, len(groups))
	for i, g := range groups {
		wg.Add(1)
		// tempGroup := GetBotsByGroupId(lpd, token, orgId, g.BotGroupID)
		go func() {
			log.Info("go concurrent get groups", "id", i)
			defer wg.Done()
			// results <- GetBotsByGroupId(lpd, token, orgId, g.BotGroupID)
			tempGroup := GetBotsByGroupId(lpd, token, orgId, g.BotGroupID)
			allBots = append(allBots, tempGroup...)
		}()
	}
	wg.Wait()
	// for each bot lookup its agent
	for i, v := range allBots {
		// TODO: this can also be concurrent
		wg.Add(1)
		go func() {
			log.Info("go concurrent get BotAgentByBotId", "id", i)
			defer wg.Done()
			botAgentsMap[v.BotID] = GetBotAgentByBotId(lpd, token, orgId, v.BotID)
		}()
	}
	// wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// cache skills
		skills := GetSkills(lpd, siteId, bearer)
		for _, s := range skills {
			skillIdToNameMap[s.ID] = s.Name
		}
	}()
	wg.Wait()

	// look up users
	users := GetUsers(lpd, siteId, bearer)
	log.Info("returned user", "count", len(users))
	for _, u := range users {
		// TODO: skillid to skill name
		// agentSkillsMap[u.LoginName] = fmt.Sprintf("%v", u.SkillIds)
		var skillName string
		if len(u.SkillIds) > 0 {
			skillName = skillIdToNameMap[u.SkillIds[0]]
		}
		agentSkillsMap[u.LoginName] = skillName
	}

	listOfBots := aggregateBots(allBots)

	elapsed := time.Since(start)
	log.Info("list of bots count", "count", len(listOfBots), "execution time", elapsed)

	// get ungroup list
	return listOfBots, nil
}

func aggregateBots(allBots []GroupBot) []Bot {
	var bots = []Bot{}
	for _, v := range allBots {
		var newBot Bot
		newBot.ID = v.BotID
		newBot.Name = v.BotName
		newBot.Group = groupNameByIdMap[v.BotGroupID]
		newBot.Agents = botAgentsMap[v.BotID]
		newBot.Skills = agentSkillsMap[botAgentsMap[v.BotID]]
		bots = append(bots, newBot)
	}

	// for each group we have to query for the bot, let's do concurrent
	// var wg sync.WaitGroup
	// results := make(chan string, len(groups))
	// log.Info("do some concurrent stuff")
	// for _, v := range groups {
	// 	wg.Add(1)
	// 	go makeRequest(v, &wg, results)
	// }
	// wg.Wait()
	// close(results)
	// log.Info("done with concurrent stuff")
	// for result := range results {
	// 	log.Info(result)
	// }
	return bots
}
