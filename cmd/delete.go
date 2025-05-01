package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove tasks from your to-do list.",
	Long:  "Deletes the specified tasks from your CSV-based to-do list, permanently removing them from the list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete Task!")
		if len(args) == 0 {
			fmt.Println("Id Not provided, try running go run . complete <id>")
			return
		}
		var updatedRecords [][]string
		id := args[0]
		f, err := os.Open("./tasks.csv")
		if err != nil {
			log.Fatal("Unable to open file!")
		}
		defer f.Close()
		r := csv.NewReader(f)
		// Used to skip the header
		records, err := r.ReadAll()
		if err != nil {
			log.Fatal("Error reading file")
		}
		found := false
		for i := 0; i < len(records); i++ {
			if records[i][0] != id {
				updatedRecords = append(updatedRecords, records[i])
			} else {
				found = true
			}
		}
		if !found {
			fmt.Println("Task not found...")
			return
		}
		// Create a loop to go through all rows in csv
		fw, err := os.OpenFile("./tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal("Unable to open file!")
		}
		defer fw.Close()
		w := csv.NewWriter(fw)
		w.WriteAll(updatedRecords)
		err = w.Error()
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
		fmt.Println("Task deleted successfully.")

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
