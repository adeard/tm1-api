package tm1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tm1-api/domain"
)

type Repository interface {
	SendTm(input domain.Tm1RequestData) (any, error)
}

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) SendTm(input domain.Tm1RequestData) (any, error) {
	var result any

	tmUrl := "http://10.120.20.174:28672/api/v1/Cubes('Running Account Master')/tm1.Update"

	inputString, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}

	var jsonStr = []byte(inputString)

	req, _ := http.NewRequest("POST", tmUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", "")

	// proxyUrl, _ := url.Parse("http://10.126.111.123:4480")
	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(response))

	return result, err
}
