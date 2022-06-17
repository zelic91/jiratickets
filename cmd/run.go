/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			panic("Please provide the file and the Jira base link")
		}
		fileName := args[0]
		jiraBaseLink := args[1]
		runMain(fileName, jiraBaseLink)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runMain(fileName string, jiraBaseLink string) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	reg := regexp.MustCompile("([a-zA-Z]+-[0-9]{2,4})")

	tickets := reg.FindAllString(string(data), -1)

	ticketLinks := map[string][]interface{}{}
	for _, t := range tickets {
		ticketLinks[fmt.Sprintf("%s/%s", jiraBaseLink, t)] = []interface{}{}
	}

	for l := range ticketLinks {
		fmt.Println(l)
	}
}
