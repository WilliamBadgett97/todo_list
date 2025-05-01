package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your todo list.",
	Long:  "Adds a new task to your CSV todo list. Provide the task description as an argument.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No task description provided!")
			return
		}

		// Open for read/write
		f, err := os.OpenFile("./tasks.csv", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal("Unable to open file:", err)
		}
		defer f.Close()

		// Check if the file is empty, write header if needed
		fileInfo, _ := f.Stat()
		if fileInfo.Size() == 0 {
			w := csv.NewWriter(f)
			w.Write([]string{"ID", "Task", "CreatedAt", "Done"})
			w.Flush()
			// Move pointer back to beginning after writing header
			f.Seek(0, 0)
		} else {
			// Move to beginning for reading
			f.Seek(0, 0)
		}

		// Read existing tasks
		r := csv.NewReader(f)
		records, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		// Determine ID
		var id int
		if len(records) == 1 && records[0][0] == "ID" {
			id = 1
		} else {
			lastID, err := strconv.Atoi(records[len(records)-1][0])
			if err != nil {
				log.Fatal(err)
			}
			id = lastID + 1
		}

		// Reposition to end before writing new entry
		f.Seek(0, io.SeekEnd)
		w := csv.NewWriter(f)

		taskName := strings.Join(args, " ")
		if strings.TrimSpace(taskName) == "" {
			fmt.Println("Task description cannot be empty.")
			return
		}

		timestamp := time.Now().Format(time.RFC3339)
		task := []string{strconv.Itoa(id), taskName, timestamp, "false"}

		if err := w.Write(task); err != nil {
			log.Fatal(err)
		}
		w.Flush()

		fmt.Printf("Task added: [%d] %s\n", id, taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
