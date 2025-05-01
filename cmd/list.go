package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all tasks in your to-do list.",
	Long:  "Displays all the tasks from your CSV-based to-do list, showing their current status and details.",
	Run: func(cmd *cobra.Command, args []string) {
		showAll, _ := cmd.Flags().GetBool("all")
		header := "ID\tTask\tCreated"
		if showAll {
			header += "\tDone"
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(w, header)
		f, err := os.Open("./tasks.csv")
		if err != nil {
			log.Fatal("Unable to open file!")
		}
		defer f.Close()
		r := csv.NewReader(f)
		_, err = r.Read()
		if err != nil {
			log.Fatal("Unable to read file!")
		}
		records, err := r.ReadAll()
		if len(records) == 0 {
			fmt.Println("No tasks found. Add one using `go run . add <task>`")
			return
		}
		for _, record := range records {
			if showAll {
				fmt.Fprintln(w, record[0]+"\t"+record[1]+"\t"+record[2]+"\t"+record[3])
			} else {
				if record[3] == "false" {
					fmt.Fprintln(w, record[0]+"\t"+record[1]+"\t"+record[2])
				}
			}
		}
		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "List all tasks")
}
