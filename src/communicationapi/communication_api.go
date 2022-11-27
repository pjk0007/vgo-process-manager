package command

type CommunicationApi struct {
	CommandArray []Command `json:"commandArray"`
}

type Command struct {
	ClassId         string          `json:"classId"`
	Id              string          `json:"id"`
	Api             string          `json:"api"`
	Payload         Payload         `json:"payload"`
	InputPortArray  InputPortArray  `json:"inputPortArray"`
	OutputPortArray OutputPortArray `json:"outputPortArray"`
}

type Payload struct {
	ClassId string `json:"classId"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	Pid     string `json:"pid"`
}

type InputPortArray struct {
}

type OutputPortArray struct {
}