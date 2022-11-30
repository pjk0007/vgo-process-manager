package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/vcanus/vgo-process-manager/src/util"
)

type Process struct {
	Id          *int64      `json:"id"`
	ParentId    *int64      `json:"parent_id"`
	Name        string      `json:"name"`
	Npath       string      `json:"npath"`
	Pid         *int        `json:"pid"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Json        interface{} `json:"json"`
}

func JsonToProcess(path string) *Process {
	jsonFile, err := os.Open(path)
	util.CheckError(err)
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	process := new(Process)
	json.Unmarshal(byteValue, &process)

	return process
}

func (process Process) ToJson(path string){
	data, _ := json.Marshal(process)
	os.WriteFile(path,data, 0644)
}
