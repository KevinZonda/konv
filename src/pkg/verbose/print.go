package verbose

import "fmt"

var VerboseMode bool = false

var _queue []string

func fmtLog(format string, args ...any) string {
	var s string
	if len(args) == 0 {
		s = format
	} else {
		s = fmt.Sprintf(format, args...)
	}
	return "[VERBOSE] " + s
}

func PrintF(format string, args ...any) {
	if VerboseMode {
		PrintQueue()
		fmt.Println(fmtLog(format, args...))
	} else {
		_queue = append(_queue, fmtLog(format, args))
	}
}

func PrintQueue() {
	for len(_queue) > 0 {
		fmt.Println(_queue[0])
		_queue = _queue[1:]
	}
}
