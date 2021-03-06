package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

// Function to get http body and decode onto a struct that represents the returned JSON
// http://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
func GetHttp(url, user, pass string, target interface{}) error {
	// Need to create an http client to set headers, otherwie this can be
	// performed in-line
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Target is autogenerated by a JSON -> GO parser
	return json.NewDecoder(resp.Body).Decode(target)
}
