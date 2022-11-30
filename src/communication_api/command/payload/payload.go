package payload

import "github.com/vcanus/vgo-process-manager/src/communication_api/command/rest"

type Payload struct {
	ClassId string `json:"classId"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	Pid     int `json:"pid"`
}


func (payload Payload) GetPayload() rest.Rest{

	payloadProcess := PayloadProcess{
		Payload: payload,
	}

	payloadMap := map[string]rest.Rest{
		"command.rest.payload.PayloadProcess":payloadProcess,
	}

	return payloadMap[payload.ClassId]
}