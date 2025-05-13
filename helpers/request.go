package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func PostTm(tmUrl string, input any) (any, error) {
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

	result := ""

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		result = "ERROR"
		err = fmt.Errorf("%v", string(response))
	} else {
		result = string(response)
		if result == "" {
			result = "SUCCESS"
		}
	}

	return result, err
}

func GetTm(tmUrl string) (any, error) {
	var result any

	req, _ := http.NewRequest("GET", tmUrl, nil)
	req.SetBasicAuth("admin", "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	response, _ := io.ReadAll(resp.Body)
	fmt.Println(string(response))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		result = "ERROR"
		err = fmt.Errorf("%v", string(response))
	} else {
		err = json.Unmarshal(response, &result)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return nil, err
		}
	}

	return result, err
}
