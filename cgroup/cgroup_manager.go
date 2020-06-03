package cgroup

import (
	"github.com/sirupsen/logrus"
	"myDocker/subsystems"
)

type Manager struct {
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *Manager {
	return &Manager{
		Path: path,
	}
}

func (manager *Manager) Apply(pid int) {
	for _, subsystem := range subsystems.SubsystemIns {
		err := subsystem.Apply(manager.Path, pid)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (manager *Manager) Set(res *subsystems.ResourceConfig) {
	for _, subsystem := range subsystems.SubsystemIns {
		err := subsystem.Set(manager.Path, res)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (manager *Manager) Remove(cgroup string) {
	for _, subsystem := range subsystems.SubsystemIns {
		err := subsystem.Remove(cgroup)
		if err != nil {
			logrus.Error(err)
		}
	}
}
