package hexutil

import (
	"strings"
)

func IsValidBtcAddress(address string) bool {
    len := len(address)
    if len < 25 {
        return false
    }
    if strings.HasPrefix(address, "1") {
        if len >= 26 && len <= 34 {
            return true
        }
    }
    if strings.HasPrefix(address, "3") && len == 34 {
        return true
    }
    if strings.HasPrefix(address, "bc1") && len > 34 {
        return true
    }

    return false
}