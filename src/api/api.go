package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vcanus/vgo-process-manager/src/config"
	"github.com/vcanus/vgo-process-manager/src/util"
)

type RegisterOutput struct{
	Server config.Server `json:"servers"`
	Id int64 `json:"id"`
}

type TopicOutput struct{
	TopicId string `json:"topic_id"`
}

func JsonToRegisterOutput(jsonByte []byte) RegisterOutput{
	var registerOutput RegisterOutput
	json.Unmarshal(jsonByte, &registerOutput)

	return registerOutput
}

func JsonToTopicOutput(jsonByte []byte) TopicOutput{
	var topicOutput TopicOutput
	json.Unmarshal(jsonByte, &topicOutput)

	return topicOutput
}

func PostRegister(processManager config.Config) RegisterOutput {
	fmt.Println(util.ToJsonString(processManager))
	reqBody := bytes.NewBufferString(util.ToJsonString(processManager))
	resp, err := http.Post(processManager.Server+"/api/process/manager", "application/json", reqBody)
	util.CheckError(err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	util.CheckError(err)
	return JsonToRegisterOutput(respBody)
}

func PutNpath(processManager config.Config) {
	client := &http.Client{}
	reqBody := bytes.NewBufferString(util.ToJsonString(processManager))
	req, _ := http.NewRequest(http.MethodPut, processManager.Server+"/api/process/manager/npath", reqBody)
	req.Header.Set("Content-Type","application/json")
	client.Do(req)
}

func PostTopic(processManager config.Config) TopicOutput {
	reqBody := bytes.NewBufferString(util.ToJsonString(processManager))
	resp, err := http.Post(processManager.Server+"/api/topic/pm", "application/json", reqBody)
	util.CheckError(err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	util.CheckError(err)
	return JsonToTopicOutput(respBody)
}