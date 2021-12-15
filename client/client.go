package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

// Client codeforces client
type Client struct {
	Jar            *cookiejar.Jar `json:"cookies"`
	Handle         string         `json:"handle"`
	HandleOrEmail  string         `json:"handle_or_email"`
	Password       string         `json:"password"`
	Ftaa           string         `json:"ftaa"`
	Bfaa           string         `json:"bfaa"`
	LastSubmission *Info          `json:"last_submission"`
	host           string
	// proxy          string
	path   string
	client *http.Client
}

// Instance global client
var Instance *Client

// Init initialize
func Init(path, host string) {
	jar, _ := cookiejar.New(nil)
	c := &Client{Jar: jar, LastSubmission: nil, path: path, host: host, client: nil}
	if err := c.load(); err != nil {
		color.Red(err.Error())
		color.Green("Create a new session in %v", path)
	}
	Proxy := http.ProxyFromEnvironment

	c.client = &http.Client{Jar: c.Jar, Transport: &http.Transport{Proxy: Proxy}}
	if err := c.save(); err != nil {
		color.Red(err.Error())
	}
	Instance = c
}

// load from path
func (c *Client) load() (err error) {
	file, err := os.Open(c.path)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, c)
}

// save file to path
func (c *Client) save() (err error) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err == nil {
		os.MkdirAll(filepath.Dir(c.path), os.ModePerm)
		err = ioutil.WriteFile(c.path, data, 0644)
	}
	if err != nil {
		color.Red("Cannot save session to %v\n%v", c.path, err.Error())
	}
	return
}
