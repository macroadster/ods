package sdk

import (
 "fmt"
 "os"
 "github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
    Use:  "odp-sdk",
    Version: version,
    Short: "odp-sdk - a simple CLI to send data",
    Long: `Open Data Platform SDK`,

    Run: func(cmd *cobra.Command, args []string) {
      if len(args)==0 {
        cmd.Help()
      }
    },
}

func init() {
  rootCmd.PersistentFlags().String("endpoint", "http://localhost:8080", "API Server URL")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}
