package systemed

// 重启系统
func Reboot() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("reboot")
	rs, err := sys.Exec()
	return rs, err
}

// 关闭系统，切断电源
func PowerOff() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("poweroff")
	rs, err := sys.Exec()
	return rs, err
}

// CPU停止工作
func Halt() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("halt")
	rs, err := sys.Exec()
	return rs, err
}

// 暂停系统
func Suspend() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("suspend")
	rs, err := sys.Exec()
	return rs, err
}

// 让系统进入冬眠状态
func Hibernate() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("hibernate")
	rs, err := sys.Exec()
	return rs, err
}

// 让系统进入交互式休眠状态
func HybridSleep() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("hybrid-sleep")
	rs, err := sys.Exec()
	return rs, err
}

// 启动进入救援状态（单用户状态）
func Rescue() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("rescue")
	rs, err := sys.Exec()
	return rs, err
}
