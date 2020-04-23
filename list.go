package systemed

// 列出您 Linux 系统上的所有单元
func ListUnitFiles() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("list-unit-files").SetArgs("--all").SetArgs("--plain").SetArgs("|grep -vE 'STATE|listed'|tr -s '\n'|awk '{print $1\" \" $2}'")
	rs, err := sys.Exec()
	return rs, err
}

// 查看当前系统的所有 Unit
func ListUnits() ([]byte, error) {
	sys := NewSystemed("systemctl")
	sys.SetArgs("list-units").SetArgs("--all").SetArgs("--plain")
	rs, err := sys.Exec()
	return rs, err
}
