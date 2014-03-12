package btsync

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var Debug bool

func NewClient(host, port, user, password string) *Client {
	return &Client{host, port, user, password}
}

// Base Url http://host:port/api
func (c Client) baseUrl() string {
	return fmt.Sprintf("http://%s:%s/api", c.Host, c.Port)
}

// Create HTTP Request and Parse JSON to struct
func (c Client) call(method string, v url.Values, r interface{}) error {
	if v == nil {
		v = make(url.Values)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", c.baseUrl(), nil)
	if err != nil {
		return err
	}

	v.Add("method", method)
	req.URL.RawQuery = v.Encode()
	req.SetBasicAuth(c.User, c.Password)

	if Debug {
		log.Printf("Request: %s\n", req.URL)
	}

	// Make Request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if Debug {
		log.Printf("Response: %s\n", data)
	}

	if err = json.Unmarshal(data, &r); err != nil {
		return err
	}

	return nil
}
