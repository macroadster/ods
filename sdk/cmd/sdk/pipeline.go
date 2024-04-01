package sdk

import (
    "fmt"

    "ds/sdk/pkg/sdk"
    "github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
    Use:   "inspect [pipeline]",
    Aliases: []string{"insp"},
    Short:  "Inspects a data pipeline",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      endpoint, _ := cmd.Flags().GetString("endpoint")
      if endpoint != "" {
        sdk.SetEndpoint(endpoint)
      }
      pipeline := args[0]
      if err := sdk.Inspect(pipeline); err != nil {
        fmt.Printf("%s\n", err)
      }
    },
}

var createCmd = &cobra.Command{
    Use:   "create [pipeline]",
    Short:  "Create a data pipeline",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      endpoint, _ := cmd.Flags().GetString("endpoint")
      if endpoint != "" {
        sdk.SetEndpoint(endpoint)
      }
      pipeline := args[0]
      if err := sdk.Create(pipeline); err != nil {
        fmt.Printf("%s\n", err)
      } else {
        fmt.Printf("Pipeline %s created.\n", pipeline)
      }
    },
}

var deleteCmd = &cobra.Command{
    Use:   "delete [pipeline]",
    Short:  "Delete a data pipeline",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      endpoint, _ := cmd.Flags().GetString("endpoint")
      if endpoint != "" {
        sdk.SetEndpoint(endpoint)
      }
      pipeline := args[0]
      if err := sdk.Delete(pipeline); err != nil {
        fmt.Printf("%s\n", err)
      } else {
        fmt.Printf("Pipeline %s deleted.\n", pipeline)
      }
    },
}

func init() {
    rootCmd.AddCommand(createCmd)
    rootCmd.AddCommand(deleteCmd)
    rootCmd.AddCommand(inspectCmd)
}
