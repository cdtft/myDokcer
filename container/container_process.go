package container

import (
	"os"
	"os/exec"
	"syscall"
)

// tty 是否创建命令行链接
func NewParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init", command}
	// /proc/self是指当前运行进程的自己环境
	// exe自己调自己，通过这种方式对创建出来的进程进行初始化
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
