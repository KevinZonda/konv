package loader

import (
	"strings"

	"github.com/KevinZonda/apt-pac/pkg/decision"
	"github.com/KevinZonda/apt-pac/pkg/utils"
)

func appendCommandToTree(s string, root *decision.CheckTree, data string) {
	ss := strings.Split(s, " ")
	var nonEmpty []string
	ss = utils.TrimAll(ss)
	for _, c := range ss {
		nonEmpty = append(nonEmpty, c)
	}
	_d := data
	root.AddToSub(nonEmpty, &_d)
}
