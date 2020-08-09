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
}

var (
    greeting string
)



func init() {
    hiCmd := &cobra.Command{
        Use: "hi --greeting 'how are you today?'",
        Short: "Say Hello!",
    Run: func(cmd *cobra.Command, args []string) {
        CoreHello()
    },
}

    rootCmd.AddCommand(hiCmd)
    hiCmd.PersistentFlags().StringVarP(&greeting, "greeting", "g", "How are you?", "Greeting statement")
}

func CoreHello() {
    fmt.Println(greeting)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
