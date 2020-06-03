package cgroup

import "myDocker/subsystems"

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

}
