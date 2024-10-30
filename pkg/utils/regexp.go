package utils

import (
	"regexp"
)

func RegexpToken(token string) (bool, error) {
	if matched, err := regexp.MatchString(`^[A-Za-z0-9/.]{243}$`, token); err != nil {
		return false, err
	} else {
		return matched, nil
	}
}
