// from https://github.com/josegm/yummy/blob/master/credentials.go

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type authResponse struct {
	Pkg          string
	Url          string
	Status       string
	Delta_ms     int
	Server       string
	Session      string
	Api_mgmt_ms  int
	Version      string
	Access_token string
}

type Credentials struct {
	ClientId     string
	ClientSecret string
	GrantType    string
	Username     string
	Password     string
	AccessToken  string
	Status       string
}

const NO_ENV_FOUND_ERROR = "You must set DELICIOUS_CLIENT_ID, DELICIOUS_CLIENT_SECRET, DELICIOUS_USERNAME AND DELICIOUS_PASSWORD as ENV variables"

func (c *Credentials) isValid() bool {
	return len(c.ClientId) > 0 && len(c.ClientSecret) > 0 && len(c.Username) > 0 && len(c.Password) > 0
}

func NewCredentials() (*Credentials, error) {
	auth := new(Credentials)
	auth.ClientId = clientID
	auth.ClientSecret = clientSecret
	auth.Username = username
	auth.Password = password
	//	env := os.Environ()
	//	for _, e := range env {
	//		key_value := strings.Split(e, "=")
	//		switch key_value[0] {
	//		case "DELICIOUS_CLIENT_ID":
	//			auth.ClientId = key_value[1]
	//		case "DELICIOUS_CLIENT_SECRET":
	//			auth.ClientSecret = key_value[1]
	//		case "DELICIOUS_USERNAME":
	//			auth.Username = key_value[1]
	//		case "DELICIOUS_PASSWORD":
	//			auth.Password = key_value[1]
	//		}
	//	}

	//	if !auth.isValid() {
	//		return nil, errors.New(NO_ENV_FOUND_ERROR)
	//	}
	return auth, nil
}

const auth_url = "https://avosapi.delicious.com/api/v1/oauth/token"

func (c *Credentials) Authenticate() error {
	values := url.Values{}
	values.Set("client_id", c.ClientId)
	values.Set("client_secret", c.ClientSecret)
	values.Set("grant_type", "credentials")
	values.Set("username", c.Username)
	values.Set("password", c.Password)

	resp, err := http.PostForm(auth_url, values)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Auth failed. Response %s.", resp.Status)
		return errors.New(msg)
	}

	defer resp.Body.Close()

	auth_response := new(authResponse)

	err = json.NewDecoder(resp.Body).Decode(&auth_response)
	if err != nil {
		return err
	}

	c.AccessToken = auth_response.Access_token
	c.Status = auth_response.Status

	return nil
}
