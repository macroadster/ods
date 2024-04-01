package sdk

import (
    "ds/sdk/pkg/sdk"
    "github.com/spf13/cobra"
    "os"
    "fmt"
)

var sendCmd = &cobra.Command{
    Use:   "send [message]",
    Short:  "Send a message",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        endpoint, _ := cmd.Flags().GetString("endpoint")
        if endpoint != "" {
          sdk.SetEndpoint(endpoint)
        }
        err := sdk.Send(args)
        if err != nil {
          fmt.Println("Error sending message: ", err)
          os.Exit(1)
        } else {
          fmt.Println("Message sent.")
        }
    },
}

func init() {
    rootCmd.AddCommand(sendCmd)
}
