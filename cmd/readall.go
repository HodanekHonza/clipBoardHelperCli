package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type ShoppingRecord struct {
	name string
	test string
}

func createShoppingList(data [][]string) []ShoppingRecord {
	var shoppingList []ShoppingRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec ShoppingRecord
			for j, field := range line {
				if j == 0 {
					rec.name = field
				} else if j == 1 {
					rec.test = field
				}
			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}

// Create a new command
var readAllCmd = &cobra.Command{
	Use:   "readall",
	Short: "Prints a greeting message",
	Long:  "A longer description that spans multiple lines and likely contains examples and usage of using this command.",
	Run: func(cmd *cobra.Command, args []string) {

		f, err := os.Open("result.csv")
		if err != nil {
			log.Fatal(err)
		}

		// remember to close the file at the end of the program
		defer f.Close()

		// read csv values using csv.Reader
		csvReader := csv.NewReader(f)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		// convert records to array of structs
		shoppingList := createShoppingList(data)

		// print the array
		fmt.Printf("%+v\n", shoppingList)

	},
}

// Initialize and add flags
func init() {
	// Add the greet command as a subcommand of the root command
	rootCmd.AddCommand(readAllCmd)
}
