package config

type ConfigJson struct {
	DeviceInfo  ConfigJsonDeviceInfo    `json:"deviceInfo"`
	ProcessInfo []ConfigJsonProcessInfo `json:"processInfo"`
}

type ConfigJsonDeviceInfo struct {
	Company   string `json:"company"`
	Site      string `json:"site"`
	Operation string `json:"operation"`
	Eq        string `json:"eq"`
}

type ConfigJsonProcessInfo struct {
	ProcessType string `json:"type"`
	ProcessPath string `json:"path"`
}
