package tm1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tm1-api/domain"
	"tm1-api/helpers"
)

type Repository interface {
	GetTm(uri1 string, uri2 string, queryString string) (any, error)
	SendTm(input domain.Tm1RequestData) (any, error)
	SendDynamicTm(input domain.Tm1DynamicRequestData) (any, error)
	SendRaTest(input domain.Tm1RequestDynamicData) (any, error)
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

func (r *repository) SendDynamicTm(input domain.Tm1DynamicRequestData) (any, error) {
	tmUrl := "http://" + input.Url + ":" + input.Port + "/api/v1/Cubes('" + input.Cubes + "')/tm1.Update"

	return helpers.PostTm(tmUrl, input.Tm1DynamicInputData)
}

func (r *repository) SendRaTest(input domain.Tm1RequestDynamicData) (any, error) {
	tmUrl := "http://10.120.20.174:25772/api/v1/Cubes('RA Test')/tm1.UpdateCells"

	return helpers.PostTm(tmUrl, input)
}

func (r *repository) GetTm(uri1 string, uri2 string, queryString string) (any, error) {

	tmUrl := "http://10.120.20.174:25772/api/v1/" + uri1 + "/" + uri2 + "?" + queryString

	req, _ := http.NewRequest("GET", tmUrl, nil)
	req.SetBasicAuth("admin", "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(response))

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return result, err
}
