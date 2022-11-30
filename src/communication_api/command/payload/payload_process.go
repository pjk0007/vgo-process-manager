package payload

import (
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
	cmd.Start()
	return cmd.Process.Pid
}

func (p PayloadProcess) Delete() any {
	process, err := os.FindProcess(p.Pid)
	if err != nil{
		return false
	}
	err = process.Kill()
	return err == nil
}