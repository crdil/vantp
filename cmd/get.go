package cmd

import (
	"fmt"
	"projects/cobratest/globals"
	"projects/cobratest/httpRequest"
	"projects/cobratest/utils"

	"github.com/spf13/cobra"
)

var headers map[string]interface{}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "GET request",
	Long:  `Make a GET request`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(globals.Headers) > 0 {
			headers = utils.StringsToHeaderMap(globals.Headers)
		}

		request := httpRequest.Request{
			Type:        "GET",
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
	rootCmd.AddCommand(getCmd)
}
