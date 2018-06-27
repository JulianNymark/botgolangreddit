package main

import (
	"encoding/json"
	"fmt"
	rq "github.com/juliannymark/botgolangreddit/http"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Printf("Hello, world.\n")

	req := rq.HttpRequest{
		Type:    "GET",
		URL:     "https://reqres.in/api/users/2",
		Headers: http.Header{},
	}

	resp, err := req.Send(nil)

	if err != nil {
		fmt.Errorf("Could not create http request")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("Could not read body")
	}

	var data map[string]interface{}
	_ = json.Unmarshal(body, &data)

	fmt.Printf("RESP:\n %v", data)

}
