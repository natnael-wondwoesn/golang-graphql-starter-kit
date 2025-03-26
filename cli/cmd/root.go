package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "go-graphqlify",
	Short:   "A CLI tool for generating Go GraphQL APIs",
	Version: version,
	Long: `GoGraphQLify is a CLI tool that helps you generate production-ready
GraphQL APIs using Go, with various database and feature options.

Complete documentation is available at https://github.com/natnael-wondwoesn/golang-graphql-starter-kit`,
	Run: func(cmd *cobra.Command, args []string) {
		displayBanner()
		fmt.Println("Use 'go-graphqlify create' to start a new project")
		fmt.Println("Use 'go-graphqlify --help' for more information")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func displayBanner() {
	banner := `
   ______      ______                 _     ______ _ _  __       
  / ____/___  / ____/________ _____  / /_  / ____/(_) |/ /_ _  __
 / / __/ __ \/ / __/ ___/ __ '/ __ \/ __ \/ / __/ / /   / / | / /
/ /_/ / /_/ / /_/ / /  / /_/ / /_/ / / / / /_/ / / /   | /| |/ / 
\____/\____/\____/_/   \__,_/ .___/_/ /_/\____/_/_/|_/|_/ |___/  
                           /_/                                   
`
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Println(green(banner))
	fmt.Printf("%s v%s - The magical GraphQL starter kit for Go\n\n", green("GoGraphQLify"), version)
}