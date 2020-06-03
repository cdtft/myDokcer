package subsystems

type CpuSetSubsystem struct {

}

func (s *CpuSetSubsystem) Set(cgroupPath string, config *ResourceConfig) error {
	return nil
}

func (s *CpuSetSubsystem) Remove(cgroupPath string) error {
	return nil
}

func (s *CpuSetSubsystem) Name() string {
	return "cpuset"
}

func (s *CpuSetSubsystem) Apply(cgroup string, pid int) error {
	return nil
}