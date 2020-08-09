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
        CoreHello()
    },
}

var (
    greeting string
)



func init() {
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
