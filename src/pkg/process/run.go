package process

import (
	"os"
	"os/exec"
)

type Runable struct {
	Name string
	Args []string
}

func (r Runable) run() {
	RunAndWait(r.Name, r.Args...)
}

func RunAndWait(name string, args ...string) {
	cmd := exec.Command(name, args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	_ = cmd.Start()
	_ = cmd.Wait()
}

func Runs(rs []Runable) {
	for _, r := range rs {
		r.run()
	}
}
