package main

import (
	"AetherGo/internal/app"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "aether",
	Short: "AetherGo is a lightweight modular web framework",
	Long:  `AetherGo is a lightweight web framework built in Go for fast and scalable applications.`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  `Create a new project with the given name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please specify the project name")
			return
		}
		projectName := args[0]
		err := app.CreateNewProject(projectName)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
