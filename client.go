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
	Username string
	Password string
}

// Locations gets information about locations for all accounts.
// https://github.com/tmaiaroto/smartthings-unofficial-docs/blob/master/Documentation.md#all-locations
func (c Client) Locations() ([]interface{}, error) {
	httpClient, err := c.newHTTPClient()
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Get(
		"https://graph-na02-useast1.api.smartthings.com/api/locations",
	)
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
	_, err = httpClient.Head(
		"https://graph-na02-useast1.api.smartthings.com/",
	)
	if err != nil {
		return nil, err
	}
	return httpClient, err
}
