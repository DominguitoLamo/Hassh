package logger

import (
	"fmt"
	"hassh/src/internal/constant"

	"github.com/rs/zerolog/log"
)

const LOGGER_LEVEL = constant.DEBUG_LEVEL

func DebugLog(format string, v... interface{}) {
	if LOGGER_LEVEL > constant.DEBUG_LEVEL {
		return
	}
	msg := fmt.Sprintf(format, v...)

    log.Debug().
        Msg(msg)
}

func ErrorLog(format string, v... interface{}) {
	if LOGGER_LEVEL > constant.ERROR_LEVEL {
		return
	}
	msg := fmt.Sprintf(format, v...)

    log.Error().
        Msg(msg)
}