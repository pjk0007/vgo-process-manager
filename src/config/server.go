package config

import "encoding/json"

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