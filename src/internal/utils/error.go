package utils

import (
	"github.com/zeromicro/x/errors"
	hconstant "hassh/src/internal/constant"
)

func ParameterError() error {
	return errors.New(hconstant.ERROR_PARAMETER, "parameter error")
}

func IpFormatError(msg string) error {
	return errors.New(hconstant.ERROR_IPFORMAT, msg)
}

func ConfigError(msg string) error {
	return errors.New(hconstant.ERROR_CONFIG, msg)
}

func SSHConnectionError(msg string) error {
	return errors.New(hconstant.ERROR_SSHCONNECTION, msg)
}

func SSHCmdError(msg string) error {
	return errors.New(hconstant.ERROR_SSHCMD, msg)
}