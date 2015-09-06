package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func deliGo(url, tags, description string) string {
	auth, err := NewCredentials()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = auth.Authenticate()
	if err != nil {
		fmt.Println(err.Error())
	}

	return AddLink(auth, url, tags, description)
}

func readDeliciousData() {
	data, err := ioutil.ReadFile("delicious.json")
	if err != nil {
		panic(err)
	}
	config := &DeliciousData{}
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}

	clientID = config.ClientID
	clientSecret = config.ClientSecret
	username = config.Username
	password = config.Password
}

var clientID string
var clientSecret string
var username string
var password string

type DeliciousData struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}
