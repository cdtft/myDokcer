package subsystems

import (
	"fmt"
	"testing"
)

func TestFindCgroupMountPoint(t *testing.T) {
 	fmt.Println(FindCgroupMountPoint("memory"))
}

func TestGetAndCreateCgroupPath(t *testing.T) {

}