package http

import (
	"fmt"
	"io"
	"net/http"
)

// HttpRequest generate a new HTTP request
type HttpRequest struct {
	URL     string
	Type    string
	Headers http.Header
}

// Send sends a http request
func (options *HttpRequest) Send(body io.Reader) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(options.Type, options.URL, body)
	if err != nil {
		return nil, err
	}

	req.Header = options.Headers
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Errorf("Request [%s] %s returned status code %v ", options.Type, options.URL, resp.StatusCode)
	}

	return resp, nil
}
