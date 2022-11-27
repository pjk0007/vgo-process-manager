package command

import (
	"os/exec"
)

type PayloadProcess struct{
	Path string
}

func (p PayloadProcess) Create() int {
	cmd := exec.Command(p.Path)
	cmd.Start()
	return cmd.Process.Pid
}