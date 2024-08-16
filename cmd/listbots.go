/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wlai-lp/bo-botflow/internal/lpapi"
)

type bot struct {
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

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}
	str := fmt.Sprintf("%d. %s", index+1, i)
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}
	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("You chose %s", m.choice))
	}
	return "\n" + m.list.View() + "\nPress q to quit.\n"
}

// listbotsCmd represents the listbots command
var listbotsCmd = &cobra.Command{
	Use:   "listbots",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Executing listbots subcommand")

		// merge commands
		viper.BindPFlags(cmd.Flags())

		bots, err := getListOfBots()
		if err != nil {
			log.Fatal("Unable to get list of bots")
		}		

		log.Info("bots", "length", len(bots))

		siteId := viper.Get("LP_SITE")
		log.Info("lp site id from viper called ", "site", siteId)

		items := []list.Item{}

		for _, b := range bots {
			items = append(items, item(b.Name))
		}
		

		const defaultWidth = 20

		l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
		l.Title = "My List"
		l.SetShowStatusBar(false)
		l.SetFilteringEnabled(false)
		l.Styles.PaginationStyle = paginationStyle
		l.Styles.HelpStyle = helpStyle

		m := model{list: l}

		p := tea.NewProgram(m)
		if err := p.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func checkListbotsConfig() error {
	// TODO: implement validation
	var valid = false
	if valid {
		fmt.Println("check list bot config return error")
		return errors.New("missing required field")
	} else {
		return nil
	}
}

func getListOfBots() ([]bot, error) {
	// get domain by siteid
	lpd, err := lpapi.GetDomain(fmt.Sprint(viper.Get("LP_SITE")))
	if err != nil {
		return nil, err
	}

	// get bot access token and orgid
	sid := fmt.Sprint(viper.Get("LP_SITE"))
	b := fmt.Sprint(viper.Get("BEARER"))
	if b == "" {
		log.Error("bearer token value is empty")
		return nil, errors.New("bearer token value is empty")
	}
	token, orgid, err := lpapi.GetBotAccessToken(lpd, b)
	if err != nil {
		return nil, err
	}
	log.Info(fmt.Sprintf("token is %v and org is %v", token, orgid))

	// get bot group list to get group id
	groups := lpapi.GetBotGroups(lpd, token, orgid)
	log.Info("total of", "groups", len(groups))

	// cache groups id to name
	// groupNameById := make(map[string]string)
	for _, g := range groups {
		groupNameByIdMap[g.BotGroupID] = g.BotGroupName
	}

	// get bots by group id
	allBots := lpapi.GetBotsByGroupId(lpd, token, orgid, UNASSIGNED)
	log.Info("total of ungroup", "ungrouped", len(allBots))

	// loop each botgroup id to get bots in the group, add it to allbots
	// todo: make this call concurrent
	for _, g := range groups {
		tempGroup := lpapi.GetBotsByGroupId(lpd, token, orgid, g.BotGroupID)
		allBots = append(allBots, tempGroup...)
	}

	// for each bot lookup its agent
	
	for _, v := range allBots {
		botAgentsMap[v.BotID] = lpapi.GetBotAgentByBotId(lpd, token, orgid, v.BotID)
	}

	// TODO: look up skill and cache it
	skills := lpapi.GetSkills(lpd, sid, b)
	for _, s := range skills {
		skillIdToNameMap[s.ID] = s.Name
	}
	

	// look up users
	users := lpapi.GetUsers(lpd, sid, b)
	log.Info("returned user", "count", len(users))
	for _, u := range users {
		// TODO: skillid to skill name		
		// agentSkillsMap[u.LoginName] = fmt.Sprintf("%v", u.SkillIds)
		var skillName string
		if len(u.SkillIds) > 0{
			skillName = skillIdToNameMap[u.SkillIds[0]]
		}
		agentSkillsMap[u.LoginName] = skillName
	}



	listOfBots := aggregateBots(allBots)
	
	log.Info("list of bots count", "count", len(listOfBots))

	// get ungroup list
	return listOfBots, nil
}

func aggregateBots(allBots []lpapi.GroupBot) []bot {
	var bots = []bot{}
	for _, v := range allBots {
		var newBot bot
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

func makeRequest(group lpapi.BotGroupsData, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	results <- fmt.Sprintf("return from make reqeust")
}

func init() {
	log.SetReportCaller(true)
	// log.WithPrefix("listbots").Info("init")
	log.Debug("init")

	LoadViperConfig()
	err := checkListbotsConfig()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// siteId := viper.Get("LP_SITE")
	// fmt.Println("xxx siteid directory:", siteId)
	rootCmd.AddCommand(listbotsCmd)
	// listbotsCmd.Flags().String("name", "", "Name to be used")
	listbotsCmd.Flags().StringVarP(&bearer, "BEARER", "B", "", "bearer token")
	// listbotsCmd.MarkFlagRequired("bearer")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listbotsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listbotsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
