package config

import (
	"encoding/json"
	"io"
	"os"
	"path"

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

	return config
}

func (config Config) ToJson(path string){
	data, _ := json.Marshal(config)
	os.WriteFile(path,data, 0644)
}

// 모든 자식 프로세스 Config를 변경한다. (서버주소, 부모Id)
func (config Config) SetProcessConfig(server Server){
	for _, processInfo := range config.Json.ProcessInfo {
		configPath := path.Dir(processInfo.ProcessPath) + "/config/"
		process := JsonToProcess(configPath + "process.config")
		process.ParentId = config.Id
		process.ToJson(configPath + "process.config")

		server.ToJson(configPath + "servers.config")
	}
}