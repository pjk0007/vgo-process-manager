package config

import (
	"encoding/json"
	"os"
)

type Server struct {
	TopicManager  string
	DeviceManager string
	Nats          []string
	Zookeeper     []string
}

func JsonToServer(jsonByte []byte) Server {
	var server Server
	json.Unmarshal(jsonByte, &server)

	return server
}

func (server Server) ToJson(path string){
	data, _ := json.Marshal(server)
	os.WriteFile(path,data, 0644)
}