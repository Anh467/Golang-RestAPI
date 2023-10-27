package main

import (
	"common"
	"encoding/json"
	"entities"

	"fmt"
)

var configPath string = "config.json"

func setConfigPath() {
	path, err := common.GetRootProject(configPath)
	if err != nil {
		fmt.Println("err: " + err.Error())
	} else {
		fmt.Println("path: " + path)
	}
	configPath = path
}

func getConfigJson(path string) entities.Config {

	file, err := common.GetFile(path)

	var config entities.Config
	if err != nil {
		fmt.Print("[main/main.go/getConfigJson(path string)]: " + err.Error())
		return config
	}

	fileContent, err := common.GetFileContents(file)

	if err != nil {
		fmt.Print("[main/main.go/getConfigJson(path string)]: " + err.Error())
		return config
	}
	errr := json.Unmarshal([]byte(fileContent), &config)
	if errr != nil {
		fmt.Print("[main/main.go/getConfigJson(path string)]: " + err.Error())
		return config
	}
	return config
}

func main() {

	setConfigPath()
	configJson := getConfigJson(configPath)
	fmt.Print(configJson)
	check := solution("abc", "bc")
	fmt.Print("Test: ", check)
}
func solution(str, ending string) bool {
	// Your code here!
	var i int
	strTh := len(str)
	endingTh := len(ending)
	if strTh < endingTh {
		return false
	}
	for i = endingTh - 1; i >= 0; i-- {

		if str[strTh-(endingTh-i)] != ending[i] {
			return false
		}
	}
	return true
}
