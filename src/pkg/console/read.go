package console

import (
	"bufio"
	"os"
	"strings"

	"github.com/KevinZonda/konv/pkg/utils"
)

func ReadYes() bool {
	scanner := bufio.NewReader(os.Stdin)
	ans, _ := scanner.ReadString('\n')
	ans = strings.ToLower(utils.Trim(ans))
	return ans == "y"
}
