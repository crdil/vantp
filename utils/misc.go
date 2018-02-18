package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"projects/cobratest/globals"
	"strings"
)

var postBody []byte
var argsBody []string
var contentType string

// StringInSlice :: used to iterate through slice and return true if arg(string) is found
func StringInSlice(findString string, findInList []string) bool {
	for _, item := range findInList {
		if item == findString {
			return true
		}
	}
	return false
}

// StringsToJSON ::
// :: takes a []string that contains elements like "key=value"
// :: and turns it into a json string by first creating a map and then marshal to json
func StringsToJSON(jsonString []string) []byte {
	var property string
	var propValue string
	var jsonmap map[string]interface{}
	jsonmap = make(map[string]interface{})

	for _, keyValue := range jsonString {
		splitData := strings.Split(keyValue, "=")
		if len(splitData)%2 != 0 {
			fmt.Println("Wrong args")
			os.Exit(1)
		}
		for index, value := range splitData {
			if index == 0 {
				property = value
			} else {
				propValue = value
			}
			jsonmap[property] = propValue
		}
	}
	jsonText, err := json.Marshal(&jsonmap)
	if err != nil {
		fmt.Println("Error")
	}
	return jsonText
}

func ParseFormdata(data []string) []byte {
	var key string
	var setValue string
	if globals.ContentType != "form" {
		fmt.Println(`Can't use "application/json" content-type with formdata`)
		os.Exit(1)
	}
	urlValues := url.Values{}
	for _, keyValue := range data {
		splitData := strings.Split(keyValue, "==")
		if len(splitData)%2 != 0 {
			fmt.Println("Wrong args")
			os.Exit(1)
		}
		for index, value := range splitData {
			if index == 0 {
				key = value
			} else {
				setValue = value
			}
			urlValues.Set(key, setValue)
		}
	}
	return []byte(urlValues.Encode())
}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func ParseArgBody(args []string) []byte {
	// if argsbody is raw json
	// if args body does not contain = assume it is the raw data to send instead
	if IsJSON(args[1]) {
		postBody = []byte(args[1])
	} else if !strings.Contains(args[1], "=") {
		postBody = []byte(args[1])
	} else if strings.Contains(args[1], "==") {
		postBody = ParseFormdata(args[1:])
	} else {
		argsBody = args[1:]
		postBody = StringsToJSON(argsBody)
	}
	return postBody
}

// ParseURL :: used to check http prefix and if not present, then validate with ParseRequestURI
func ParseURL(sentURL string) string {
	// hasHTTPprefix := strings.HasPrefix(sentURL, `http://`)
	var newURL string
	if strings.HasPrefix(sentURL, `http://`) {
		newURL = sentURL
	} else if strings.HasPrefix(sentURL, `https://`) {
		newURL = sentURL
	} else {
		newURL = `http://` + sentURL
	}
	_, err := url.ParseRequestURI(newURL)
	if err != nil {
		log.Fatalln("Invalid URL")
	}
	return newURL
}

func StringsToHeaderMap(headers []string) map[string]interface{} {
	var property string
	var propValue string
	// store the headers
	var mapObj map[string]interface{}

	mapObj = make(map[string]interface{})

	for _, headerFlags := range headers {
		splitData := strings.Split(headerFlags, ":")
		if len(splitData)%2 != 0 {
			fmt.Println("Invalid header")
			os.Exit(1)
		}
		for index, value := range splitData {
			if index == 0 {
				property = value
			} else {
				propValue = value
			}
			mapObj[property] = propValue
		}
	}
	return mapObj
}

func ParseBasicAuth(authCredentials string) []string {
	if !strings.Contains(authCredentials, ":") {
		return []string{}
	}

	credentials := strings.Split(authCredentials, ":")

	if len(credentials) > 2 {
		fmt.Println("Error in authCredentials")
	}
	return credentials
}
