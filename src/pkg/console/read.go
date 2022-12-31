package console

import (
	"bufio"
	"github.com/KevinZonda/konv/pkg/utils"
	"os"
    "strings"
)

func ReadYes() bool {
	scanner := bufio.NewReader(os.Stdin)
	ans, _ := scanner.ReadString('\n')
    ans = strings.ToLower(utils.Trim(ans))
	return ans == "y"
}
