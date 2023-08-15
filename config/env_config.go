package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var jsonFile *string = nil

func init() {
	//jsonFile = flag.String("env", "dev", "a string")
	//flag.Parse()
	jsonFile = flag.String("file", "", "a String")
}

type Server struct {
	Address string `json:"address"`
}

type Context struct {
	Timeout int `json:"timeout"`
}

type Configuration struct {
	Debug   bool    `json:"debug"`
	Server  Server  `json:"server"`
	Context Context `json:"context"`
}

var (
	configuration *Configuration = nil
)

func LoadAppConfiguration(fileName string) *Configuration {
	if fileName != " " && configuration == nil {
		jsonFile, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error in reading the file = %s", err)
			return nil
		}

		err = json.Unmarshal(jsonFile, &configuration)
		if err != nil {
			fmt.Printf("Error in unmarshalling = %s", err)
			return nil
		}

	}
	return configuration
}
