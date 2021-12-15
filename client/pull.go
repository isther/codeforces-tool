package client

import (
	"errors"
	"regexp"
)

func findMessage(body []byte) (string, error) {
	reg := regexp.MustCompile(`Codeforces.showMessage\("([^"]*)"\);\s*?Codeforces\.reformatTimes\(\);`)
	tmp := reg.FindSubmatch(body)
	if tmp != nil {
		return string(tmp[1]), nil
	}
	return "", errors.New("Cannot find any message")
}
