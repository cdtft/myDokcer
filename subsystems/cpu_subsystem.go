package subsystems

type CpuSubsystem struct {

}

func (s *CpuSubsystem) Set(cgroupPath string, config *ResourceConfig) error {
	return nil
}

func (s *CpuSubsystem) Remove(cgroupPath string) error {
	return nil
}

func (s *CpuSubsystem) Name() string {
	return "cpu"
}

func (s *CpuSubsystem) Apply(cgroup string, pid int) error {
	return nil
}