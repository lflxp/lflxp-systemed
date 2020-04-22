package systemed

// 显示当前主机的信息
func Hostnamectl() ([]byte, error) {
	sys := NewSystemed("hostnamectl")
	rs, err := sys.Exec()
	return rs, err
}

// 查看本地化设置
func Localectl() ([]byte, error) {
	sys := NewSystemed("localectl")
	rs, err := sys.Exec()
	return rs, err
}

// 查看当前时区设置
func Timedatectl() ([]byte, error) {
	sys := NewSystemed("timedatectl")
	rs, err := sys.Exec()
	return rs, err
}

// 查看当前登录的用户
func Loginctl() ([]byte, error) {
	sys := NewSystemed("loginctl")
	rs, err := sys.Exec()
	return rs, err
}
