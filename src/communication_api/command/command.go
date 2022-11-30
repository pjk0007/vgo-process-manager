package command

import (
	"github.com/vcanus/vgo-process-manager/src/communication_api/command/payload"
	"github.com/vcanus/vgo-process-manager/src/communication_api/command/portarray"
)

type Command struct {
	ClassId         string          `json:"classId"`
	Id              string          `json:"id"`
	Api             string          `json:"api"`
	Payload         payload.Payload         `json:"payload"`
	InputPortArray  portarray.InputPortArray  `json:"inputPortArray"`
	OutputPortArray portarray.OutputPortArray `json:"outputPortArray"`
}

func (command Command) Execute(){
	rest := command.Payload.GetPayload()
	if command.ClassId == "command.rest.Create" {
		rest.Create()
	}else if command.ClassId == "command.rest.Delete"{
		rest.Delete()
	}
}