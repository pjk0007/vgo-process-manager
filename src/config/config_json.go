package config

type ConfigJson struct {
	EquipmentInfo  ConfigJsonEquipmentInfo    `json:"equipmentInfo"`
	ProcessInfo []ConfigJsonProcessInfo `json:"processInfo"`
}

type ConfigJsonEquipmentInfo struct {
	Company   string `json:"company"`
	Site      string `json:"site"`
	Operation string `json:"operation"`
	Eq        string `json:"eq"`
}

type ConfigJsonProcessInfo struct {
	ProcessType string `json:"type"`
	ProcessPath string `json:"path"`
}
