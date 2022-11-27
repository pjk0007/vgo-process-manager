package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/vcanus/vgo-process-manager/src/util"
)

type Config struct {
	Id          *int64 `json:"id"`
	Pid         int	`json:"pid"`
	Name        string `json:"name"`
	Npath       string `json:"npath"`
	Json        ConfigJson `json:"json"`
	Description string `json:"description"`
	Server string `json:"server"`
}

func JsonToConfig(path string) *Config{
	jsonFile, err := os.Open(path)
	util.CheckError(err)
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var config = new(Config)
	json.Unmarshal(byteValue, &config)

	// if config.Id == 0{
	// 	config.Id = nil
	// }
	return config
}

