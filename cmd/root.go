package cmd
import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  homedir "github.com/mitchellh/go-homedir"

)


var (
    cfgFile  string
    greeting string
)

var rootCmd = &cobra.Command{
  Use:   "higo",
  Short: "Higo is a cobra & viper variable input demo",
  Long: "",
  //	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default: $PWD/hello.yaml)")
  rootCmd.PersistentFlags().StringVarP(&greeting, "greeting", "g", "How are you?", "Greeting statement")
}


func initConfig() {
  if cfgFile != "" {
    viper.SetConfigFile(cfgFile)
  } else {
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    viper.AddConfigPath(home)
    viper.SetConfigName("hello.yaml")
  }

  viper.AutomaticEnv() // read in environment variables that match

  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}
