package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config Serv
type Config struct {
	Port string `json:"port"`
	Root string `json:"root"`
}

// LoadConfig read the jByte and returns Config struct
func LoadConfig() *Config {

	config := &Config{
		Port: "8080",
		Root: "./www/",
	}

	jByte, err := ioutil.ReadFile("./config.json")
	if err != nil {
		initConfig(config)
		return config
	}

	jByte = stirpBOM(jByte)
	jsonErr := json.Unmarshal(jByte, &config)
	if jsonErr != nil {
		fmt.Println("cant parse json: ", jsonErr.Error())
		return config
	}

	return config
}

func stirpBOM(fileBytes []byte) []byte {
	trimmedBytes := bytes.Trim(fileBytes, "\xef\xbb\xbf")
	return trimmedBytes
}

func initConfig(config *Config) {
	// Create config.json
	jByte, mErr := json.MarshalIndent(config, "", "  ")
	if mErr == nil {
		ioutil.WriteFile("./config.json", jByte, 0777)
	}

	// Create www folder
	_, err := os.Stat("./www")
	if err != nil {
		if e := os.Mkdir("./www", 0777); e == nil {
			greeting := "<body>Serving!</body>"
			ioutil.WriteFile("./www/index.html", []byte(greeting), 0777)
		}

	}
}
