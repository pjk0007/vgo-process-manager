package payload

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type PayloadProcess struct{
	Payload
}

func (p PayloadProcess) Create() any {

	var cmd *exec.Cmd
	if runtime.GOOS == "darwin"{
		cmd = exec.Command("open", p.Path)
	} else if runtime.GOOS == "windows"{
		cmd = exec.Command(p.Path)
	} else {
		cmd = exec.Command(p.Path)
	}
	cmd.Run()
	fmt.Println(cmd.Process.Pid)
	return cmd.Process.Pid
}

// Process Kill
// TODO: Process가 Kill되지 않고 Process를 감싸는 cmd, bash가 Kill되는 문제
func (p PayloadProcess) Delete() any {
	process, err := os.FindProcess(p.Pid)
	if err != nil{
		return false
	}
	err = process.Kill()
	return err == nil
}