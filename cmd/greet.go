package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "encoding/csv"
    "log"
    "os"

)

// Create a new command
var greetCmd = &cobra.Command{
    Use:   "greet",
    Short: "Prints a greeting message",
    Long:  "A longer description that spans multiple lines and likely contains examples and usage of using this command.",
    Run: func(cmd *cobra.Command, args []string) {
        name, _ := cmd.Flags().GetString("name")
        data := [][]string{
          {"originalName", "World"},
          {"newName", name},
        }
        if name == "" {
          name = "World"
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

    // Add the greet command as a subcommand of the root command
    rootCmd.AddCommand(greetCmd)
}

