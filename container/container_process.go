package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// tty 是否创建命令行链接
func NewParentProcess(tty bool) (*exec.Cmd, *os.File) {
	readPip, writePip, err := NewPipe()
	if err != nil {
		logrus.Error(err)
		return nil, nil
	}
	logrus.Info("clone namespace")
	//一开始我就有一个疑问，RunContainerInitProcess这个方法是在什么时候调用的
	// /proc/self/exe init /bash/bin ==>> /myDocker/self/exe init /bash/bin
	// [change]不再使用agrs传递command参数，而是采用优雅的管道通信的方式
	args := []string{"init"}
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
	//为什么是extraFile,因为一个进程默认会有3个文件描述分别是标准输入，标准输出，标准错误
	//在外带这个文件描述符
	cmd.ExtraFiles = []*os.File{readPip}
	mntURL := "/root/mnt/"
	rootURL := "/root/"
	NewWorkSpace(rootURL, mntURL)
	cmd.Dir = mntURL
	return cmd, writePip
}

func NewPipe() (read *os.File, write *os.File, err error) {
	r, w, e := os.Pipe()
	if e != nil {
		return nil, nil, e
	}
	return r, w, e
}

func RunContainerInitProcess() error {
	cmdArray := readUserCommand()
	if cmdArray == nil || len(cmdArray) == 0 {
		return fmt.Errorf("user command is nil")
	}

	setMount()

	//在path中寻找命令的绝对路径
	path, err := exec.LookPath(cmdArray[0])
	if err != nil {
		logrus.Error(err)
		return err
	}

	//syscall.Exec完成初始化动作并将用户进程运行起来
	if err := syscall.Exec(path, cmdArray[0:], os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}

func readUserCommand() []string {
	//uintptr(3)就是指的index为3的文件描述符，也就是额外的那个
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := ioutil.ReadAll(pipe)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	commands := string(msg)
	logrus.Info("子进程收到的命令：", commands)
	return strings.Split(commands, " ")
}

//pivot_root是一个系统调用，主要功能是改变当前的root文件系统。
//pivot_root将当前的进程的root文件系统移动到put_old文件夹中，然后使用new_root成为新的root文件系统.
func pivotRoot(root string) error {
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("mount rootfs to itself error: %v", err)
	}
	oldRootDir := filepath.Join(root, ".pivot_root")
	if err := os.Mkdir(oldRootDir, 0777); err != nil {
		return err
	}
	if err := syscall.PivotRoot(root, oldRootDir); err != nil {
		return err
	}
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chanage work space to / error : %v", err)
	}
	pivotDir := filepath.Join("/", ".pivot_root")
	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		return err
	}
	return os.Remove(pivotDir)
}

func setMount() {
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Errorf("get current location error %v", err)
		return
	}

	pivotRoot(pwd)

	//MS_NOEXEC在本文件系统中不允许运行其他程序
	//MS_NOSUID在本系统中运行程序的时候，不允许set-user-ID或set-group-ID
	//MS_NODEV所有mount的系统都会默认设定的参数RunContainerInitProcess
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
}