package process

import (
	"fmt"
	"github.com/KevinZonda/konv/pkg/utils"
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

	err := cmd.Start()
	utils.PanicIfNotNil(err,
		fmt.Sprintf(
			"run failed:\n"+
				"cmd: %s %+v\n"+
				"err: %+v\n", name, args, err))
	_ = cmd.Wait()
}

func Runs(rs []Runable) {
	for _, r := range rs {
		r.run()
	}
}
