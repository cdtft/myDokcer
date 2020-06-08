package main

import (
	log "github.com/sirupsen/logrus"
	"myDocker/cgroup"
	"myDocker/container"
	"myDocker/subsystems"
	"os"
)

func Run(tty bool, command string, res * subsystems.ResourceConfig) {
	parent, writePip := container.NewParentProcess(tty)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	cgroupManager := cgroup.NewCgroupManager("cdtftcontainer-cgroup")
	defer cgroupManager.Remove()
	cgroupManager.Set(res)
	//将进程加入cgroup
	cgroupManager.Apply(parent.Process.Pid)
	sendInitCommand(command, writePip)
	parent.Wait()
	os.Exit(-1)
}

//管道通信
func sendInitCommand(command string, writePip *os.File) {
	log.Infof("init command is:[%s]", command)
	writePip.WriteString(command)
	writePip.Close()
}
