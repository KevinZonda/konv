package process

import (
	"os"
	"os/exec"
)

type Runable struct {
	Name string
	Args []string
}

func (r Runable) run() error {
	return RunAndWait(r.Name, r.Args...)
}

func RunAndWait(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func Runs(rs []Runable, ifErr func(Runable, error) bool) {
	for _, r := range rs {
		err := r.run()
		if err != nil && ifErr != nil {
			if ifErr(r, err) {
				break
			}
		}
	}
}
