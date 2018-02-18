package globals

// TODO: refactor, should not use global vars

// ContentType :: used to set the content type for the request
var ContentType string

// Headers :: used to set headers for request
var Headers []string

// Timeout :: set the timeout for the request
var Timeout int

// NoSSLVerify :: disable SSL verification
var NoSSLVerify bool

// BasicAuth :: basic authentication with request
var BasicAuth string
