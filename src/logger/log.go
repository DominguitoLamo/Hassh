package logger

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func DebugLog(format string, v... interface{}) {
	msg := fmt.Sprintf(format, v...)

    log.Debug().
        Msg(msg)
}

func ErrorLog(format string, v... interface{}) {
	msg := fmt.Sprintf(format, v...)

    log.Error().
        Msg(msg)
}