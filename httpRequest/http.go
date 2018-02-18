package httpRequest

// http://polyglot.ninja/golang-making-http-requests/

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"projects/cobratest/utils"
	"regexp"
	"time"
)

var (
	jsonRegex = `application/(.*)json`
)

// Request : used when calling makeRequest function
type Request struct {
	Type        string
	URL         string
	Data        []byte
	Headers     map[string]interface{}
	Timeout     int
	NoSSLVerify bool
	BasicAuth   []string
	FormData    []string
}

// MakeRequest :: used to make the HTTP request TODO: refactor into separate http methods for each request type, check golang request for inspiration
func MakeRequest(reqObj *Request, verboseRequest bool, contentType string) string {
	requestURL := utils.ParseURL(reqObj.URL)
	var (
		requestHeaders  string
		responseHeaders string
		client          *http.Client
		request         *http.Request
		// requestContent  io.Reader
	)

	// if post body is empty
	if len(reqObj.Data) == 0 {
		reqObj.Data = nil
	}

	// set the timeout using the timeout flag
	timeout := time.Duration(reqObj.Timeout) * time.Millisecond
	// if no ssl verification
	if reqObj.NoSSLVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr, Timeout: timeout}
	} else {
		client = &http.Client{Timeout: timeout}
	}

	request, err := http.NewRequest(reqObj.Type, requestURL, bytes.NewBuffer(reqObj.Data))

	request.Header.Set("User-Agent", "Vantp/0.0.1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set request content type
	if contentType == "json" {
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")
	} else if contentType == "text" {
		request.Header.Set("Content-Type", "text/html")
	} else if contentType == "form" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// set basic authetication
	if len(reqObj.BasicAuth) > 0 {
		username := reqObj.BasicAuth[0]
		password := reqObj.BasicAuth[1]
		request.SetBasicAuth(username, password)
	}

	// set request headers using global flag header
	if len(reqObj.Headers) > 0 {
		for key, value := range reqObj.Headers {
			request.Header.Set(key, value.(string))
		}
	}

	// if verboseRequest use httputils DumpRequestOut to print headers
	if verboseRequest {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		requestHeaders = string(dump)
	}
	// print request headers and body, verbose request
	if len(requestHeaders) > 0 {
		fmt.Println(requestHeaders)
		fmt.Println()
	}

	// send request
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// close body
	defer resp.Body.Close()

	// always show response headers
	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	responseHeaders = string(dump)
	if len(responseHeaders) > 0 {
		fmt.Print(responseHeaders)
	}

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	contenttype := resp.Header.Get(`Content-type`)

	// if json content type in response header, prettify output
	match, err := regexp.MatchString(jsonRegex, contenttype)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if match {
		colorResponse, err := utils.Format(body)
		if err != nil {
			fmt.Println(err)
		}

		return string(colorResponse)
	}

	// colorize text response body
	return utils.FormatText(body)
}

// CheckErr :: TODO: refactor
func CheckErr(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}
