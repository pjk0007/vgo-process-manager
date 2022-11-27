package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/nats-io/nats.go"
	"github.com/vcanus/vgo-process-manager/src/api"
	command "github.com/vcanus/vgo-process-manager/src/communicationapi"
	"github.com/vcanus/vgo-process-manager/src/config"
)

func main(){
	fmt.Println("Hello World!");

	// ProcessManager Config
	var ProcessManager = config.JsonToConfig("config.json");
	ProcessManager.Pid = os.Getpid()

	// Register ProcessManager
	// - get process id
	// - get servers (topic manager, device manager, nats, zookeeper)
	registerOutput := api.PostRegister(*ProcessManager)
	ProcessManager.Id = &registerOutput.Id
	var Servers = registerOutput.Server

	// Zookeeper Config
	Zookeeper, _, _ := zk.Connect(Servers.Zookeeper,time.Second)
	ProcessManager.Npath, _ = Zookeeper.Create(fmt.Sprintf("/process/%d",*ProcessManager.Id),[]byte{},zk.FlagEphemeral,zk.WorldACL(zk.PermAll))

	// Put Npath to Device manager
	api.PutNpath(*ProcessManager)

	// Get Command Topic (from Topic manager)
	var CommandTopic = api.PostTopic(*ProcessManager).TopicId
	fmt.Println(CommandTopic)

	// Nats Connect
	// - connect nats 
	// - subscribe command topic
	// - execute command when command comes in
	nc, _ := nats.Connect(strings.Join(Servers.Nats, ", "))
	natsConn, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	natsConn.Subscribe(CommandTopic, func(cmd string){
		fmt.Println(cmd)
		command.ExecuteAll(cmd)
	})


	/////////////// TEST /////////////////////

	/*
	 * process start (This should be in command)
	 */
	var cmd = exec.Command("./process/edgev1/edgev.exe")
	cmd.Run()
	// var pid = cmd.Process.Pid
}