package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
)

// tty 是否创建命令行链接
func NewParentProcess(tty bool, command string) *exec.Cmd {
	logrus.Info("clone namespace")
	//一开始我就有一个疑问，RunContainerInitProcess这个方法是在什么时候调用的
	// /proc/self/exe init /bash/bin ==>> /myDocker/self/exe init /bash/bin
	args := []string{"init", command}
	// /proc/self是指当前运行进程的自己环境
	// exe自己调自己，通过这种方式对创建出来的进程进行初始化,这里的自己是指的myDocker这个程序
	cmd := exec.Command("/proc/self/exe", args...)
	//克隆namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWUTS,
	}

	//打开命令行连接
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}

func RunContainerInitProcess(command string, args []string) error {
	//MS_NOEXEC在本文件系统中不允许运行其他程序
	//MS_NOSUID在本系统中运行程序的时候，不允许set-user-ID或set-group-ID
	//MS_NODEV所有mount的系统都会默认设定的参数RunContainerInitProcess
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	logrus.Infof("runContainerInitProcess command %s", command)
	//syscall.Exec完成初始化动作并将用户进程运行起来
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}
