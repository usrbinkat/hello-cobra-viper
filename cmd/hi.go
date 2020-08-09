package cmd

import (
    "os"
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
    hiCmd.PersistentFlags().StringVarP(&greeting "greeting", "g", "How are you?", "Greeting statement")
}

func CoreHello() {

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
