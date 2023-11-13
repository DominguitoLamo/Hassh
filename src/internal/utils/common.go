package utils

import (
	"strconv"
	"strings"
)

func IpFormatValid(ip string) error {
	numbers := strings.Split(ip, ".")
	if len(numbers) != 4 {
		return IpFormatError("This is not IPV4 format!!!!")
	}

	for _, n := range numbers {
		i, err := strconv.Atoi(n)
		if err != nil {
			return IpFormatError("Part of the ip is not number")
		}

		if i < 0 || i > 255 {
			return IpFormatError("Part of the ip number is not in the range from 0 to 255")
		}
	}

	return nil
}