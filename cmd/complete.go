package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark tasks as completed in your to-do list.",
	Long:  "Marks all the tasks as completed and updates your CSV-based to-do list accordingly.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Id Not provided, try running go run . complete <id>")
			return
		}
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
		for i := 1; i < len(records); i++ {
			if records[i][0] == id {
				records[i][3] = "true" // Mark as done
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Task not found")
		}
		// Create a loop to go through all rows in csv
		fw, err := os.OpenFile("./tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal("Unable to open file!")
		}
		defer fw.Close()
		w := csv.NewWriter(fw)
		w.WriteAll(records)
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
