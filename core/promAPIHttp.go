package core

import (
	"encoding/json"
	"fmt"
	"goPrometheus/models"
	"io/ioutil"
	"net/http"
	"time"
)

var responsePayloadList models.ResponsePayloadList

// var responseRulesList models.RulesList
// var responseAlerts models.Alerts

var method string
var urlFull string

func ConnectHttp(apiUrlHttp string) {

	start := time.Now()
	method = "GET"
	startTime := "2022-01-01T00:01:00.000Z"
	endTime := "2022-02-01T00:01:00.000Z"
	durationTime := "&" + startTime + "&" + endTime
	query := "query?query=sum%28rate%28ops_response_latency_ms_bucket%7Ble%3D%221000%22%2C+direction%3D%22inbound%22%2C+namespace%3D%22sre%22%2C++status%21%7E%225..%22%7D%5B5m%5D%29%29+by+%28alias%29+%2F++ignoring%28le%29++sum%28rate%28ops_response_latency_ms_count%7Bdirection%3D%22inbound%22%2C+namespace%3D%22sre%22%7D%5B5m%5D%29%29+by+%28alias%29" + durationTime
	urlFull = apiUrlHttp + query

	client := &http.Client{}
	req, err := http.NewRequest(method, urlFull, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	json.Unmarshal(bodyBytes, &responsePayloadList)
	prettyJson, _ := json.MarshalIndent(responsePayloadList, "", "   ")

	fmt.Println(string(prettyJson))
	fmt.Println("Query Duration:", time.Since(start))

}
