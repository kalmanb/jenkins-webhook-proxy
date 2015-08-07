package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const examplePayload string = `
	{
	  "actor": "User",
	  "repository": "Repository",
	  "push": {
			"changes": [
				{
					"new": {
						"type": "branch",
						"name": "name-of-branch",
						"target": {
							"type": "commit",
							"hash": "d8a044e",
							"author": "...",
							"message": "new commit message\n",
							"date": "2015-06-09T03:34:49+00:00",
							"parents": [
								{
									"hash": "1e65c05c1d5171631d92438a13901ca7dae9618c",
									"type": "commit"
								}
							]
						}
					},
					"old": {
						"type": "branch",
						"name": "name-of-branch",
						"target": {
							"type": "commit",
							"hash": "1e65c05c1d5171631d92438a13901ca7dae9618c",
							"author": "...",
							"message": "old commit message\n",
							"date": "2015-06-08T21:34:56+00:00",
							"parents": [
								{
									"hash": "e0d0c2041e09746be5ce4b55067d5a8e3098c843",
									"type": "commit"
								}
							]
						}
					},
					"created": false,
					"forced": false,
					"closed": false
				}
			]
		}
	}`

func TestJsonParsing(t *testing.T) {
	assert.Equal(t, getCommit([]byte(examplePayload)), "d8a044e")
}

func TestIntHttp(t *testing.T) {
	url := "http://localhost:8080/commitPushed"

	jsonStr := []byte(examplePayload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	assert.Equal(t, getCommit([]byte(examplePayload)), "d8a044e")
}
