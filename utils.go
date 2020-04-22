package systemed

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func ExecCommand(cmd string) ([]byte, error) {
	pipeline := exec.Command("/bin/sh", "-c", cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	pipeline.Stdout = &out
	pipeline.Stderr = &stderr
	err := pipeline.Run()
	if err != nil {
		return stderr.Bytes(), err
	}
	// fmt.Println(stderr.String())
	return out.Bytes(), nil
}

type Systemed struct {
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
	Result  string   `json:"result"`
	Command string   `json:"command"`
}

func NewSystemed(cmd string) *Systemed {
	return &Systemed{
		Cmd: cmd,
	}
}

// 合并命令
func (this *Systemed) GetCommand() error {
	if this.Cmd == "" {
		return errors.New("cmd is none")
	}

	this.Command = fmt.Sprintf("%s %s", this.Cmd, strings.Join(this.Args, " "))
	return nil
}

func (this *Systemed) SetArgs(args ...string) *Systemed {
	this.Args = append(this.Args, args...)
	return this
}

// help命令解释
func (this *Systemed) Help() ([]byte, error) {
	s, e := ExecCommand("nmap -h")
	return s, e
}

// 执行scan命令并获取结果
func (this *Systemed) Exec() ([]byte, error) {
	var (
		result []byte
		err    error
	)
	// 获取命令
	this.GetCommand()
	// 执行命令
	// fmt.Println(this.Command)
	result, err = ExecCommand(this.Command)

	return result, err
}
