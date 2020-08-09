package cmd
import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  homedir "github.com/mitchellh/go-homedir"

)


var cfgFile string

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
  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "config file (default: $PWD/hello.yaml)")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".go-hello-viper" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".go-hello-viper")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

