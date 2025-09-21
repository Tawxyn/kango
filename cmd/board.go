package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

type board struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Columns   string `json:"column"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type columns struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Cards []cards `json:"cards"`
	Order string  `json:"order"`
}

type cards struct {
	ID    string `json:"name"`
	Title string `json:"title"`
}

// boardCmd represents the board command
var boardCmd = &cobra.Command{
	Use:   "",
	Short: "List latest board",
	Long:  `List latest board added with all cards displayed.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("board called")
	},
}

func init() {
	rootCmd.AddCommand(boardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// boardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// boardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
