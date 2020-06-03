package subsystems

type MemorySubsystem struct {
}

func (s *MemorySubsystem) Set(cgroupPath string, config *ResourceConfig) error {
	return nil
}

func (s *MemorySubsystem) Remove(cgroupPath string) error {
	return nil
}

func (s *MemorySubsystem) Name() string {
	return "memory"
}

func (s *MemorySubsystem) Apply(cgroup string, pid int) error {
	return nil
}
