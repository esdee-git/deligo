//from https://github.com/josegm/yummy/blob/master/posts.go

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const addPostsUrl = "https://api.del.icio.us/v1/posts/add?"
const getPostsUrl = "https://api.del.icio.us/v1/posts/get?"

func AddLink(c *Credentials, urll, tags, description string) string {
	log.Println("AddLink called")
	values := url.Values{}
	values.Set("url", urll)
	values.Set("description", description)
	values.Set("tags", tags)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	req, _ := http.NewRequest("GET", addPostsUrl+values.Encode(), nil)
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}

	if res.StatusCode != 200 {
		msg := fmt.Sprintf("Auth failed. Response %s.", res.Status)
		return msg
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	bodyString := string(body)
	//delicious api currently has a flaw in returning error even when link was added
	//need to check if this is the case
	if strings.Contains(bodyString, "error adding link") {
		if strings.Contains(GetPost(c, urll), "no bookmarks") {
			return string(body)
		} else {
			return "<?xml version=\"1.0\" encoding=\"UTF-8\"?><result code=\"done\"/>"
		}
	}

	return string(body)
}

func GetPost(c *Credentials, urll string) string {
	log.Println("GetPost called")
	values := url.Values{}
	values.Set("url", urll)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	req, _ := http.NewRequest("GET", getPostsUrl+values.Encode(), nil)
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}

	if res.StatusCode != 200 {
		msg := fmt.Sprintf("Auth failed. Response %s.", res.Status)
		return msg
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	return string(body)
}
