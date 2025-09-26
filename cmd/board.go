package cmd

import (
	//"encoding/json"
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
	Use:   "boards",
	Short: "List latest board",
	Long:  `List latest board added with all cards displayed.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("board called")
	},
}

func getUserInput() {
	var input string
	fmt.Scanln(&input)
}

func init() {
	rootCmd.AddCommand(boardCmd)

}
