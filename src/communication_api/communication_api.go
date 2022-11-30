package communication_api

import (
	"encoding/json"

	"github.com/vcanus/vgo-process-manager/src/communication_api/command"
)

type CommunicationApi struct {
	CommandArray []command.Command `json:"commandArray"`
}

func ExecuteAll(commandJson []byte) {
	var communicationapi = new(CommunicationApi)
	json.Unmarshal(commandJson, &communicationapi)
	for _, command := range communicationapi.CommandArray{
		command.Execute()
	}
}