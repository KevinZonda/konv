package process

import (
	"os"
	"os/exec"
)

type Runable struct {
	Name string
	Args []string
	Env  []string
}

func (r Runable) run() error {
	cmd := exec.Command(r.Name, r.Args...)

	cmd.Env = r.Env
	
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
