package cmd

import (
	"fmt"
	"projects/cobratest/globals"
	"projects/cobratest/httpRequest"
	"projects/cobratest/utils"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "DELETE request",
	Long:  `Make a HTTP DELETE request`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(globals.Headers) > 0 {
			headers = utils.StringsToHeaderMap(globals.Headers)
		}

		request := httpRequest.Request{
			Type:        "DELETE",
			URL:         args[0],
			Data:        nil,
			Headers:     headers,
			Timeout:     globals.Timeout,
			NoSSLVerify: globals.NoSSLVerify,
			BasicAuth:   utils.ParseBasicAuth(globals.BasicAuth),
		}
		body := httpRequest.MakeRequest(&request, Verbose, globals.ContentType)
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
