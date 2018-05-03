package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Db     string
	Port   int
	Secret string
}

var Config *config

func init() {
	file, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println(err)
		return
	}
}
