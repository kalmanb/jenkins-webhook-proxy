package jenkins

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func StartJenkinsBuild(commit string) {
	host := "jenkins.movio.co"
	job := "mm-kalstest-deleteme"
	token := "0MnMBtOmOIIv"

	url := fmt.Sprintf("http://%s/buildByToken/buildWithParameters?job=%s&token=%s&GitCommit=%s", host, job, token, commit)

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
}
