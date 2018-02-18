package cmd

import (
	"fmt"
	"projects/cobratest/globals"
	"projects/cobratest/httpRequest"
	"projects/cobratest/utils"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "PUT request",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(globals.Headers) > 0 {
			headers = utils.StringsToHeaderMap(globals.Headers)
		}

		argURL := args[0]
		postBody := utils.ParseArgBody(args)

		request := httpRequest.Request{
			Type:        "PUT",
			URL:         argURL,
			Data:        postBody,
			Headers:     headers,
			Timeout:     globals.Timeout,
			NoSSLVerify: globals.NoSSLVerify,
			BasicAuth:   utils.ParseBasicAuth(globals.BasicAuth),
		}
		body := string(httpRequest.MakeRequest(&request, Verbose, globals.ContentType))
		fmt.Println(body)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
