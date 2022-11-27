package util

import (
	"encoding/json"
)

func ToJsonString(config any) string {
	jsonByte, err := json.Marshal(config)
	CheckError(err)
	return string(jsonByte)
}