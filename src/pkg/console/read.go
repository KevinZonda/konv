package console

import (
	"bufio"
	"github.com/KevinZonda/konv/pkg/utils"
	"os"
)

func ReadYes() bool {
	scanner := bufio.NewReader(os.Stdin)
	ans, _ := scanner.ReadString('\n')
	return utils.Trim(ans) != "y"
}
