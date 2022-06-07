//go:build linux
package osname

import (
	"errors"
	"io/ioutil"
	"regexp"
)

func osname() (string, error) {
	data, e := ioutil.ReadFile("/etc/os-release")
	if e != nil {
		return "", wrapErr(e, "Read os-release failed")
	}
	m := regexp.MustCompile("PRETTY_NAME\\s*=\\s*(?:\\\"([^\\\"\\n]+)\\\"|([^\\n]+))").FindStringSubmatch(string(data))
	if m == nil {
		return "", errors.New("Parse os-release failed")
	}
	if m[1] == "" {
		return m[2], nil
	}
	return m[1], nil
}
