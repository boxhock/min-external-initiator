package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Service struct {
	ChainlinkNode ChainlinkNode
}

func main() {
	clNodeUrl, err := url.Parse(os.Getenv("CHAINLINK_URL"))
	if err != nil {
		panic(err)
	}
	clAccess := os.Getenv("CHAINLINK_ACCESS")
	clSecret := os.Getenv("CHAINLINK_SECRET")

	clNode := ChainlinkNode{
		Endpoint:     *clNodeUrl,
		AccessKey:    clAccess,
		AccessSecret: clSecret,
	}

	service := &Service{ChainlinkNode: clNode}
	RunWebserver(service)
}

type JobConfig struct {
	Url string `json:"url"`
	Method string `json:"method"`
}

func (s Service) SubscribeToJob(jobid string, conf JobConfig) {
	go func() {
		resp, err := triggerUrl(conf.Url, conf.Method)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Initial response:", string(resp))

		t := time.NewTicker(1 * time.Minute)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				newResp, err := triggerUrl(conf.Url, conf.Method)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println("New response:", string(newResp))

				if bytes.Equal(resp, newResp) {
					// Responses are equal, do not report
					continue
				}

				// Responses were different, send job trigger
				err = s.ChainlinkNode.triggerJob(jobid)
				if err != nil {
					fmt.Println(err)
					continue
				}

				resp = newResp
			}
		}
	}()
}

func triggerUrl(url, method string) ([]byte, error) {
	fmt.Println("Sending", method, "trigger to URL", url)

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode < 200 || r.StatusCode >= 400 {
		return nil, errors.New("got unexpected status code")
	}

	return ioutil.ReadAll(r.Body)
}
