package btsync

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"regexp"
	"net/url"
	"encoding/json"
)

// BTSync Client
type Client struct {
	Host     string
	Port     string
	User     string
	Password string
	token    string
	cookies  *myjar
	values   url.Values
}

// Secret
type Secret struct {
	Secret   string   `json:"secret"`
	ROSecret string   `json:"rosecret"`
}

// AddSyncFolder or RemoveSyncFolder Result
type SyncFolderError struct {
	Error   int			`json:"error"`
	Message string		`json:"message"`
	Path    string		`json:"n"`
	Secret
	Err     error
}

// Folder from GetSyncFolders
type Folder struct {
	Name   string  	`json:"name"`
	Secret string	`json:"secret"`
	Size   string	`json:"size"`
}

type FolderInfo struct {
	Folders  []Folder `json:"folders"`
	Speed	string   `json:"speed"`
}

// Get http://host:port/gui/ url
func (c *Client) BaseUrl() string {
	return fmt.Sprintf("http://%s:%s/gui/", c.Host, c.Port)
}

// Create new Request with cookies and authorization
func (c *Client) NewRequest() (*http.Request, *http.Client, error) {
	client := &http.Client{}

	if c.cookies == nil {
		c.cookies = &myjar{}
		c.cookies.jar = make(map[string] []*http.Cookie)
		c.values = make(url.Values)
	}

	client.Jar = c.cookies
	req, err := http.NewRequest("GET", c.BaseUrl(), nil)
	if err != nil {
		return nil, nil, err
	}

	c.values.Set("token", c.token)
	req.URL.RawQuery = c.values.Encode()
	req.SetBasicAuth(c.User, c.Password)

	return req, client, err
}

// Fill c.token and return operation result
func (c *Client) RequestToken() bool {
	req, client, err := c.NewRequest();
	if err != nil {
		log.Print(err)
		return false
	}

	req.URL.Path += "token.html"
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error : %s", err)
		return false
	}

	data, _ := ioutil.ReadAll(resp.Body)
	c.token = regexp.MustCompile(`</?[^>]+>`).ReplaceAllString(string(data), "")

	return c.token != ""
}

// Get Secret
func (c *Client) GenerateSecret() ( s Secret, err error ) {
	c.values.Set("action", "generatesecret")

	req, client, err := c.NewRequest();
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(data , &s); err != nil {
		return
	}

	return
}

// Get Sync Folders
func (c *Client) Folders() ( fi FolderInfo, err error ) {
	c.values.Set("action", "getsyncfolders")
	req, client, err := c.NewRequest();
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(data , &fi); err != nil {
		return
	}

	return
}

// Add Sync Folder
func (c *Client) AddSyncFolder(path, secret string) ( r SyncFolderError ) {
	if secret == "" {
		s, err := c.GenerateSecret()
		if err != nil {
			r.Err = err
			return
		}

		secret = s.Secret
		r.Secret = s
	} else {
		r.Secret.Secret = secret
	}

	c.values.Set("action", "addsyncfolder")
	c.values.Set("name", path)
	c.values.Set("secret", secret)

	req, client, err := c.NewRequest();
	if err != nil {
		r.Err = err
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		r.Err = err
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if r.Err = json.Unmarshal(data , &r); r.Err != nil {
		return
	}

	if r.Error != 0 {
		r.ROSecret = ""
	}

	return
}

// Remove Sync Folder
func (c *Client) RemoveSyncFolder(path, secret string) ( r SyncFolderError ) {
	c.values.Set("action", "removefolder")
	c.values.Set("name", path)
	c.values.Set("secret", secret)

	req, client, err := c.NewRequest();
	if err != nil {
		r.Err = err
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		r.Err = err
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if r.Err = json.Unmarshal(data , &r); r.Err != nil {
		return
	}

	return
}

