package subsystems

type ResourceConfig struct {
	MemoryLimit string //内存限制
	CpuShare    string //CPU时间片权重
	CpuSet      string //CPU核心数
}

type Subsystem interface {
	Name() string                               //return subsystem name
	Set(path string, res *ResourceConfig) error //给cgroup设置在这个subsystem中的资源限制
	Apply(path string, pid int) error           //add process to cgroup
	Remove(path string) error                   //remove some one cgroup limit
}

var SubsystemIns = []Subsystem{
	&CpuSetSubsystem{},
	&CpuSubsystem{},
	&MemorySubsystem{},
}
