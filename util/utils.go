package util

import (
    "github.com/charmbracelet/log"
    "fmt"
	"runtime"
)

func getCaller() string {
    pc, _, _, ok := runtime.Caller(2)
    if !ok {
        return "unknown"
    }
    fn := runtime.FuncForPC(pc)
    if fn == nil {
        return "unknown"
    }
    return fn.Name()
}

// Helper function to log and return an error
func LogAndReturnError(msg string) error {
	msg = fmt.Sprintf("%v caller=%v", msg, getCaller())
    log.Error("An error occurred", "err", msg)	
    return fmt.Errorf(msg)
}

