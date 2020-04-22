package systemed

// 查看启动时的默认 Target
func GetDefault(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("get-default").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 设置启动时的默认 Target
func SetDefault(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("set-default").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// # 切换 Target 时，默认不关闭前一个 Target 启动的进程，
// # systemctl isolate 命令改变这种行为，
// # 关闭前一个 Target 里面所有不属于后一个 Target 的进程
func Isolate(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("isolate").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}
