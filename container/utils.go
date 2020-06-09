package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

//对容器和镜像进一步的隔离
func NewWorkSpace(rootURL string, mntURL string) {
	createReadOnlyLayer(rootURL)
	createWriteLayer(rootURL)
	createMountPoint(rootURL, mntURL)
}

func DeleteWorkSpace(rootURL string, mntURL string) {
	deleteMountPoint(mntURL)
	deleteWriteLayer(rootURL)
}

//解压busybox.tar作为容器的只读层
func createReadOnlyLayer(rootURL string) {
	busyboxURL := rootURL + "busybox/"
	busyboxTarUrl := rootURL + "busybox.tar"
	//如果不存在
	if !pathExist(busyboxURL) {
		if err := os.Mkdir(busyboxURL, 0777); err != nil {
			logrus.Errorf("mkdir dir %s error. %v", busyboxURL, err)
		}
		//解压busybox.tar
		if _, err := exec.Command("tar", "-xvf", busyboxTarUrl, "-C", busyboxURL).CombinedOutput(); err != nil {
			logrus.Errorf("解压busybox.tar失败 error: %v", err)
		}
	}
}

func createWriteLayer(rootURL string) {
	writeURL := rootURL + "writeLayer/"
	if err := os.Mkdir(writeURL, 0777); err != nil {
		logrus.Errorf("mkdir dir:%s err: %v", writeURL, err)
	}
}

func deleteWriteLayer(rootURL string) {
	writeLayerURL := rootURL + "writeLayer/"
	if err := os.RemoveAll(writeLayerURL); err != nil {
		logrus.Errorf("delete write layer path: %s, error: %v", rootURL, err)
	}
}

func createMountPoint(rootURL string, mntURL string) {
	//创建挂在目录
	if err := os.Mkdir(mntURL, 0777); err != nil {
		logrus.Errorf("mkdir dir %s err: %v", mntURL, err)
		return
	}
	//将writeLayer和readOnlyLayer挂在到mntURL下
	dirs := "dirs=" + rootURL + "writeLayer:" + rootURL + "busybox/"
	cmd := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", mntURL)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		logrus.Errorf("exec mount error: %v", err)
	}
}

func deleteMountPoint(mntURL string) {
	cmd := exec.Command("umount", mntURL)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		logrus.Errorf("umount mount point error: %v", err)
	}
	if err := os.RemoveAll(mntURL); err != nil {
		logrus.Errorf("remove mount path: %s, error:%v", mntURL, err)
	}
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		if os.IsNotExist(err) {
			return true
		}
		return false
	}
}
