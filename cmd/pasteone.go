package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// Create a new command
var greetCmd = &cobra.Command{
	Use:   "pasteone",
	Short: "Prints a greeting message",
	Long:  "A longer description that spans multiple lines and likely contains examples and usage of using this command.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		content, _ := cmd.Flags().GetString("content")
		data := [][]string{
			{"Name", "Content"},
			{name, content},
		}
		if content == "" {
			content = "Default Value"
		}
		if name == "" {
			name = "Default Value"
		}
		// create a file
		file, err := os.Create("result.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// initialize csv writer
		writer := csv.NewWriter(file)

		defer writer.Flush()

		// write all rows at once
		writer.WriteAll(data)

		fmt.Printf("Hello, %s!\n", name)
	},
}

// Initialize and add flags
func init() {
	// Add a flag to the greet command
	greetCmd.Flags().StringP("name", "n", "", "Name to greet")
	greetCmd.Flags().StringP("content", "c", "", "Content to store")

	// Add the greet command as a subcommand of the root command
	rootCmd.AddCommand(greetCmd)
}
