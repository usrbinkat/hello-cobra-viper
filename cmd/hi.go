package cmd

import (
//	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var hiCmd = &cobra.Command{
	Use:   "hi",
	Short: "Cobra & Viper Golang Hello Demo",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var (
	greeting string
)

func init() {
	rootCmd.AddCommand(hiCmd)
	hiCmd.PersistentFlags().StringP(&greeting "greeting", "g", "How are you?", "Greeting statement")
}

func CoreHello() {
    //
    fmt.Println("hi called")
}
