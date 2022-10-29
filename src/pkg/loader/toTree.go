package loader

import (
	"strings"

	"github.com/KevinZonda/apt-pac/pkg/decision"
)

func parseCommandToTree(s string) *decision.CheckTree {
	ss := strings.Split(s, " ")
	root := &decision.CheckTree{}
	curr := root
	for i, c := range ss {
		if i == 0 {
			root.Key = c
			continue
		}
		curr.AddKey(c)
		curr = root.Children[c]
	}
	return root
}
