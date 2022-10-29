package loader

import (
	"strings"

	"github.com/KevinZonda/apt-pac/pkg/decision"
	"github.com/KevinZonda/apt-pac/pkg/utils"
)

func appendCommandToTree(s string, root *decision.CheckTree, data string) {
	ss := strings.Split(s, " ")
	ss = utils.TrimAll(ss)
	root.AddToSub(ss, data)
}
