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
	req.Header.Add("Content-Type", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	// executing request
	client := &http.Client{}
	response, err := client.Do(req)
	checkErr(err)

	// receiving response
	respBody, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(respBody))

	// parsing response
	var cwData interface{}
	json.Unmarshal(respBody, &cwData) //extracting the json file
	//fmt.Println(cwData)

	// taking required values
	var goScore float64 = -1

	ranks, isOK := cwData.(map[string]interface{})["ranks"]
	if isOK {
		languages := ranks.(map[string]interface{})["languages"]
		golang := languages.(map[string]interface{})["go"]
		goScore = golang.(map[string]interface{})["score"].(float64)
	}
	//fmt.Println(goScore)

	return int(goScore)
}
