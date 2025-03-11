package sys11dbaassdk

import (
	"log"
	"net/http"
	"net/http/httputil"
	"slices"
	"strings"
)

// List of headers that need to be redacted
var REDACT_HEADERS = []string{"authorization", "x-s11-api-key"}

type requestLoggingTransport struct{}

func (t *requestLoggingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	// Get redacted headers
	orginalHeaders := request.Header
	request.Header = redactHeaders(request.Header.Clone())

	dump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		log.Printf("cannot dump outgoing request: %s", err)
	} else {
		var logString string
		logString += "[DEBUG] ********** sys11dbaas SDK Request **********\n"
		logString += "%s\n"
		logString += "[DEBUG] ********************************************\n"

		log.Printf(logString, dump)
	}

	// Send the request to the next transport
	request.Header = orginalHeaders
	response, requestError := http.DefaultTransport.RoundTrip(request)

	dump, err = httputil.DumpResponse(response, true)
	if err != nil {
		log.Printf("cannot dump ingoing response: %s", err)
	} else {
		var logString string
		logString += "[DEBUG] ********** sys11dbaas SDK Request **********\n"
		logString += "%s\n"
		logString += "[DEBUG] ********************************************\n"

		log.Printf(logString, dump)
	}

	return response, requestError
}

func redactHeaders(headers http.Header) http.Header {
	for name := range headers {
		if slices.Contains(REDACT_HEADERS, strings.ToLower(name)) {
			headers.Set(name, "REDACATED")
		}
	}

	return headers
}
