/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
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

func showTable(bots []bot) {

	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Bane", Width: 50},
		{Title: "Group", Width: 10},
		{Title: "Agent", Width: 10},
		{Title: "Skill", Width: 10},
	}

	rows := []table.Row{
		{"1", "Tokyolkfjlsak;jfl;ks djf lsd fkjhs ", "Japan", "37,274,000", "123"},
		{"2", "Delhi", "India", "32,065,760", "1234"},		
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
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

		bots, err := getListOfBots()
		if err != nil {
			log.Fatal("Unable to get list of bots")
		}		

		log.Info("retrieved all bots", "length", len(bots))

		siteId := viper.Get("LP_SITE")
		log.Info("lp site id from viper called ", "site", siteId)
		

		showTable(bots)
	},
}

func init() {
	rootCmd.AddCommand(botstableCmd)
	botstableCmd.Flags().StringVarP(&bearer, "BEARER", "B", "", "bearer token")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// botstableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// botstableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
