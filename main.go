package main

import (
	"encoding/json"
	"fmt"
	"goPrometheus/core"
	"goPrometheus/models"
	"os"
)

func main() {

	_, apiUrlCLI := InitEnv()

	core.ConnectCli(apiUrlCLI)
	// core.ConnectHttp(apiUrlHttp)

}

func InitEnv() (apiUrlHttp string, apiUrlCLI string) {

	file, errFile := os.Open("./config/config.json")
	if errFile != nil {
		os.Exit(0)
	}

	decoder := json.NewDecoder(file)
	configuration := models.Configuration{}
	errDecoder := decoder.Decode(&configuration)
	if errDecoder != nil {
		fmt.Println("error:", errDecoder)
	}
	apiUrlHttp = configuration.ApiUrlHttp
	apiUrlCLI = configuration.ApiUrlCLI

	return apiUrlHttp, apiUrlCLI

}
