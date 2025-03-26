package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/natnael_wondwoesn/GGStarter/cli/internal/tui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new Go GraphQL API project",
	Long: `Create a new Go GraphQL API project with various options.
You will be guided through an interactive UI to configure your project.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		projectPath, _ := filepath.Abs(projectName)

		// Check if directory already exists
		if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("directory %s already exists", projectPath))
		}

		// Display banner
		displayBanner()
		cyan := color.New(color.FgCyan).SprintFunc()
		fmt.Printf("%s Let's build something awesome together!\n\n", cyan("🧙‍♂️ Welcome to GoGraphQLify!"))

		// Run the Bubbletea UI
		config, err := tui.StartProjectCreator(projectName)
		if err != nil {
			return err
		}

		// If user cancelled
		if config == nil {
			return errors.New("project creation cancelled")
		}

		// Show generating message
		yellow := color.New(color.FgYellow).SprintFunc()
		fmt.Printf("\n%s Crafting your GraphQL API...\n", yellow("🚧"))

		// Generate the project using the config
		err = tui.GenerateProject(projectPath, config)
		if err != nil {
			return err
		}

		// Show success message
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s Done! Your project is ready at %s\n\n", green("✅"), projectPath)

		// Show next steps
		fmt.Println("Next steps:")
		fmt.Println("  1. cd", projectName)
		fmt.Println("  2. go mod download")
		fmt.Println("  3. make run")
		fmt.Printf("\nVisit http://localhost:8080/playground to start exploring your GraphQL API!\n")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}