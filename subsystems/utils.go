package subsystems

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

//找到某个subsystem的hierarchy cgroup根节点所在的目录
func FindCgroupMountPoint(subsystemType string) string {
	mountInfoFile, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		logrus.Warnf("打开/proc/self/mountinfo失败: %s", err)
		return ""
	}
	defer mountInfoFile.Close()
	scanner := bufio.NewScanner(mountInfoFile)
	for scanner.Scan() {
		// 44 32 0:39 / /sys/fs/cgroup/cpu,cpuacct rw,nosuid,nodev,noexec,relatime shared:22 - cgroup cgroup rw,cpu,cpuacct
		lineInfo := scanner.Text()
		fields := strings.Split(lineInfo, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystemType {
				return fields[4]
			}
		}
	}
	logrus.Infof("没有找到对应subsystem[%s]类型的cgroup路径", subsystemType)
	return ""
}

//创建cgroupPath的绝对路径，并返回
func GetAndCreateCgroupPath(subsystemType string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupMountPoint := FindCgroupMountPoint(subsystemType)
	 _, err := os.Stat(path.Join(cgroupMountPoint, cgroupPath))
	 if err != nil {
	 	if os.IsNotExist(err) && autoCreate {
	 		err := os.Mkdir(path.Join(cgroupMountPoint, cgroupPath), 0755)
	 		if err != nil {
	 			return "", err
			}
		} else {
			return "", err
		}
	 }
	 return path.Join(cgroupMountPoint, cgroupPath), nil
}
