package utils

import (
	"github.com/zeromicro/x/errors"
	hconstant "hassh/src/internal/constant"
)

func ParameterError() error {
	return errors.New(hconstant.ERROR_PARAMETER, "parameter error")
}