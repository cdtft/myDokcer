package subsystems

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubsystem struct {
}

//将内存资源限制写入memory.limit_in_bytes文件
func (s *MemorySubsystem) Set(cgroupPath string, config *ResourceConfig) error {
	absolutePath, err := GetAndCreateCgroupPath(s.Name(), cgroupPath, true)
	if err != nil {
		return err
	}
	if config.MemoryLimit != "" {
		err := ioutil.WriteFile(path.Join(absolutePath, "memory.limit_in_bytes"), []byte(config.MemoryLimit), 06444)
		if err != nil {
			return nil
		}
	}
	return nil
}

func (s *MemorySubsystem) Remove(cgroupPath string) error {
	absolutePath, err := GetAndCreateCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	os.RemoveAll(absolutePath)
	logrus.Infof("移除cgroup:[%s]", cgroupPath)
	return nil
}

func (s *MemorySubsystem) Name() string {
	return "memory"
}

func (s *MemorySubsystem) Apply(cgroupPath string, pid int) error {
	absolutePath, err := GetAndCreateCgroupPath(s.Name(), cgroupPath, true)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(absolutePath, "tasks"), []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return err
	}

	logrus.Infof("进程：%d加入cgroup:[%s]", pid, cgroupPath)
	return nil
}
