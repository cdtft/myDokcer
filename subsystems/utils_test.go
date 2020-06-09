package subsystems

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFindCgroupMountPoint(t *testing.T) {
 	fmt.Println(FindCgroupMountPoint("memory"))
}

func TestGetAndCreateCgroupPath(t *testing.T) {
	fmt.Println(filepath.Join("root", "busybox"))
}