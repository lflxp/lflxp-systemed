package systemed

// 查看启动耗时
func Analyze() ([]byte, error) {
	sys := NewSystemed("systemd-analyze")
	rs, err := sys.Exec()
	return rs, err
}

// 查看每个服务的启动耗时
func Blame() ([]byte, error) {
	sys := NewSystemed("systemd-analyze")
	sys.SetArgs("blame")
	rs, err := sys.Exec()
	return rs, err
}

// 显示瀑布状的启动过程流
func CriticalChain() ([]byte, error) {
	sys := NewSystemed("systemd-analyze")
	sys.SetArgs("critical-chain")
	rs, err := sys.Exec()
	return rs, err
}

// 显示指定服务的启动流
func CriticalChainService(service string) ([]byte, error) {
	sys := NewSystemed("systemd-analyze")
	sys.SetArgs("critical-chain").SetArgs(service)
	rs, err := sys.Exec()
	return rs, err
}
