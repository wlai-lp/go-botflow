/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wlai-lp/bo-botflow/internal/lpapi"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var baseStyle = lipgloss.NewStyle().
	// Height(20).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

type botModel struct {
	table table.Model
}

func (m botModel) Init() tea.Cmd { return nil }

func (m botModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m botModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func showTable(bots []lpapi.Bot) {

	columns := []table.Column{
		{Title: "Id", Width: 40},
		{Title: "Name", Width: 50},
		{Title: "Group", Width: 20},
		{Title: "Agent", Width: 20},
		{Title: "Skill", Width: 30},
	}

	rows := []table.Row{}

	for _, b := range bots {
		var s [5]string
		s[0] = b.ID
		s[1] = b.Name
		s[2] = b.Group
		s[3] = b.Agents
		s[4] = b.Skills
		slice := s[:]
		rows = append(rows, slice)	
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := botModel{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// botstableCmd represents the botstable command
var botstableCmd = &cobra.Command{
	Use:   "botstable",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("botstable called")

		viper.BindPFlags(cmd.Flags())
		// TODO: validate required parameters

		// get paramers
		siteId := fmt.Sprint(viper.Get("LP_SITE"))
		bearer := fmt.Sprint(viper.Get("BEARER"))
		log.Info("get env params", "site", siteId, "bearer", bearer)

		bots, err := lpapi.GetListOfBots(siteId, bearer)
		if err != nil {
			log.Fatal("Unable to get list of bots")
		}		

		log.Info("retrieved all bots", "length", len(bots))

		showTable(bots)
	},
}

func init() {
	rootCmd.AddCommand(botstableCmd)
	botstableCmd.Flags().StringVarP(&bearer, "BEARER", "B", "", "bearer token")
	botstableCmd.Flags().StringVarP(&account, "LP_SITE", "a", "", "LP account number")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// botstableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// botstableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
