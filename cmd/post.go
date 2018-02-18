package cmd

import (
	"fmt"
	"projects/cobratest/globals"
	"projects/cobratest/httpRequest"
	"projects/cobratest/utils"

	"github.com/spf13/cobra"
)

// should be a global variable since it is going to be used for all reuqests
var contentType string
var formData []string

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "POST request",
	Long:  `Post Command`,
	Args:  cobra.MinimumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		if len(globals.Headers) > 0 {
			headers = utils.StringsToHeaderMap(globals.Headers)
		}

		argURL := args[0]

		requestBody := utils.ParseArgBody(args)

		request := httpRequest.Request{
			Type:        "POST",
			URL:         argURL,
			Data:        requestBody,
			Headers:     headers,
			Timeout:     globals.Timeout,
			NoSSLVerify: globals.NoSSLVerify,
			BasicAuth:   utils.ParseBasicAuth(globals.BasicAuth),
			FormData:    formData,
		}
		body := string(httpRequest.MakeRequest(&request, Verbose, globals.ContentType))
		fmt.Println(body)
	},
}

func init() {
	postCmd.Flags().StringArrayVar(&formData, "form-data", []string{}, "Form data to send with request")
	rootCmd.AddCommand(postCmd)
}
