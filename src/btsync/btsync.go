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

// Get Params
type params map[string]string

// Get http://host:port/gui/ url
func (c *Client) BaseUrl() string {
	return fmt.Sprintf("http://%s:%s/gui/", c.Host, c.Port)
}

// Returns new Client
func NewClient(host, port, user, password string) *Client {
	return &Client{
		Host: host,
		Port: port,
		User: user,
		Password: password,
		cookies: &myjar{ make(map[string] []*http.Cookie) },
		values: make(url.Values),
	};
}

// Create new Request with cookies and authorization
func (c *Client) NewRequest(afterUrl string, getParams params) ( resp *http.Response, err error ) {
	client := &http.Client{}

	client.Jar = c.cookies
	req, err := http.NewRequest("GET", c.BaseUrl(), nil)
	if err != nil {
		return nil, err
	}

	if afterUrl != "" {
		req.URL.Path += afterUrl
	}

	// set get params
	c.values.Set("token", c.token)
	for k, v := range getParams {
		c.values.Set(k , v)
	}

	req.URL.RawQuery = c.values.Encode()
	req.SetBasicAuth(c.User, c.Password)

	// Make Request
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	return
}

// Fill c.token and return operation result
func (c *Client) RequestToken() bool {
	resp, err := c.NewRequest("token.html", nil);
	if err != nil {
		log.Print(err)
		return false
	}

	data, _ := ioutil.ReadAll(resp.Body)
	c.token = regexp.MustCompile(`</?[^>]+>`).ReplaceAllString(string(data), "")

	return c.token != ""
}

// Get Secret
func (c *Client) GenerateSecret() ( s Secret, err error ) {
	getParams := params {
		"action": "generatesecret",
	}

	resp, err := c.NewRequest("", getParams)
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
	getParams := params {
		"action": "getsyncfolders",
	}

	resp, err := c.NewRequest("", getParams);
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

	getParams := params {
		"action": "addsyncfolder",
		"name" : path,
		"secret": secret,
	}

	resp, err := c.NewRequest("", getParams);
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
	getParams := params {
		"action": "removefolder",
		"name": path,
		"secret": secret,
	}

	resp, err := c.NewRequest("", getParams);
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

