package smartthings

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// A Client is a Smartthings API Client.
type Client struct {
	Domain     string
	Username   string
	Password   string
	httpClient *http.Client
}

// Locations gets information about locations for all accounts.
// https://github.com/tmaiaroto/smartthings-unofficial-docs/blob/master/Documentation.md#all-locations
func (c Client) Locations() ([]interface{}, error) {
	resp, err := c.get("/api/locations")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data []interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Mode return current mode name of first location
func (c Client) Mode() (string, error) {
	locations, err := c.Locations()
	if err != nil {
		return "", err
	}
	location := locations[0].(map[string]interface{})
	mode := location["mode"].(map[string]interface{})
	return mode["name"].(string), nil
}

func (c Client) get(path string) (*http.Response, error) {
	if c.httpClient == nil {
		var err error
		c.httpClient, err = c.newHTTPClient()
		if err != nil {
			return nil, err
		}
	}

	url := url.URL{
		Scheme: "https",
		Host:   c.Domain,
		Path:   path,
	}
	return c.httpClient.Get(url.String())
}

func (c Client) newHTTPClient() (*http.Client, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Jar: cookieJar,
	}
	_, err = httpClient.PostForm(
		"https://auth-global.api.smartthings.com/sso/authenticate",
		url.Values{"username": {c.Username}, "password": {c.Password}},
	)
	if err != nil {
		return nil, err
	}
	url := url.URL{
		Scheme: "https",
		Host:   c.Domain,
	}
	_, err = httpClient.Head(url.String())
	if err != nil {
		return nil, err
	}
	return httpClient, nil
}
