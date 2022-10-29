package loader

import (
	"strings"

	"github.com/KevinZonda/apt-pac/pkg/decision"
	"github.com/KevinZonda/apt-pac/pkg/utils"
)

func appendCommandToTree(s string, root *decision.CheckTree, data *string) {
	ss := strings.Split(s, " ")
	var nonEmpty []string
	for _, c := range ss {
		t := utils.Trim(c)
		if t != "" {
			nonEmpty = append(nonEmpty, c)
		}
	}
	root.AddToSub(nonEmpty, data)
}
