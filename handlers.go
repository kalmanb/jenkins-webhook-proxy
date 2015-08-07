package main

import (
	"encoding/json"
	"fmt"
	"github.com/kalmanb/jenkins-webhook-proxy/jenkins"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func getCommit(jsonStr []byte) string {
	type Target struct {
		Hash string
	}
	type New struct {
		Target Target
	}
	type Change struct {
		New New
	}
	type Push struct {
		Changes []Change
	}
	type Update struct {
		Push Push
	}

	var update Update
	err := json.Unmarshal(jsonStr, &update)
	if err != nil {
		panic(err)
	}

	return update.Push.Changes[0].New.Target.Hash
}

func CommitPushed(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	commit := getCommit(body)
	log.Println("Just testing:", commit)
	jenkins.StartJenkinsBuild(commit)

	// Return ok to bitbucket
	w.WriteHeader(http.StatusOK)
}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
*/
