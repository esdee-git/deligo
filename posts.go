//from https://github.com/josegm/yummy/blob/master/posts.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const posts_url = "https://api.del.icio.us/v1/posts/add?"

func AddLink(c *Credentials, urll, tags, description string) string {
	log.Println("AddLink called")
	values := url.Values{}
	values.Set("url", urll)
	values.Set("description", description)
	values.Set("tags", tags)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", posts_url+values.Encode(), nil)
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
