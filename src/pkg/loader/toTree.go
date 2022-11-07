package loader

import (
	"strings"

	"github.com/KevinZonda/konv/pkg/decision"
	"github.com/KevinZonda/konv/pkg/utils"
)

func appendCommandToTree(s string, root *decision.CheckTree, data string) {
	ss := strings.Split(s, " ")
	ss = utils.TrimAll(ss)
	root.AddToSub(ss, data)
}
