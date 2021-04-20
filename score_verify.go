package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// codewars function takes the username of a Codewars user and returns the Go-Score (problem solved with golang) of that user
func codewars(username string) int {
	apiURL := fmt.Sprintf("https://www.codewars.com/api/v1/users/%s", username)

	//setting up new request
	req, err := http.NewRequest("GET", apiURL, nil)
	checkErr(err)

	// executing request & receiving response
	client := &http.Client{}
	response, err := client.Do(req)
	checkErr(err)
	defer response.Body.Close()

	// taking response body
	respBody, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	//fmt.Println(string(respBody))

	// parsing response
	var cwData interface{}
	json.Unmarshal(respBody, &cwData) //extracting the json file
	//fmt.Println(cwData)

	// taking required values
	var goScore float64 = -1

	ranks, isOK1 := cwData.(map[string]interface{})["ranks"]
	if isOK1 {
		languages, isOK2 := ranks.(map[string]interface{})["languages"]
		if isOK2 {
			golang, isOK3 := languages.(map[string]interface{})["go"]
			if isOK3 {
				goScore = golang.(map[string]interface{})["score"].(float64)
			}
		}
	}
	//fmt.Println(goScore)

	return int(goScore)
}
