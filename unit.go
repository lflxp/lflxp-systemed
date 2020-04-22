package systemed

// 显示单个 Unit 的状态
func Status(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("status").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 显示某个 Unit 是否正在运行
func IsActive(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("is-active").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 显示某个 Unit 是否处于启动失败状态
func IsFailed(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("is-failed").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 显示某个 Unit 服务是否建立了启动链接
func IsEnabled(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("is-enabled").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 立即启动一个服务
func Start(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("start").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 立即停止一个服务
func Stop(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("stop").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 重启一个服务
func ReStart(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("restart").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 杀死一个服务的所有子进程
func Kill(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("kill").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 重新加载一个服务的配置文件
func Reload(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("reload").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 重载所有修改过的配置文件
func DaemonReload() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("daemon-reload")
	rs, err := sys.Exec()
	return rs, err
}

// 显示某个 Unit 的所有底层参数
func Show(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("show").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 显示某个 Unit 的指定属性的值
// $ systemctl show -p CPUShares httpd.service

// # 设置某个 Unit 的指定属性
// $ sudo systemctl set-property httpd.service CPUShares=500

// 依赖关系
func ListDependencies(name string, all bool) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("list-dependencies").SetArgs(name)
	if all {
		sys.SetArgs("--all")
	}
	rs, err := sys.Exec()
	return rs, err
}

// 用于在上面两个目录之间，建立符号链接关系
// $ sudo systemctl enable clamd@scan.service
// # 等同于
// $ sudo ln -s '/usr/lib/systemd/system/clamd@scan.service' '/etc/systemd/system/multi-user.target.wants/clamd@scan.service'
func Enable(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("enable").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 撤销符号链接关系，相当于撤销开机启动
func Disable(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("disable").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}

// 查看配置文件的内容
func Cat(name string) ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("cat").SetArgs(name)
	rs, err := sys.Exec()
	return rs, err
}
