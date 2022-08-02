//go:build linux

package osname

import (
	"io/ioutil"
	"regexp"
)

func osname() (string, error) {
	data, e := ioutil.ReadFile("/etc/os-release")
	if e != nil {
		return "", _ReadOsReleaseFailed.Cause(e)
	}
	m := regexp.MustCompile("PRETTY_NAME\\s*=\\s*(?:\\\"([^\\\"\\n]+)\\\"|([^\\n]+))").FindStringSubmatch(string(data))
	if m == nil {
		return "", _ReadOsReleaseFailed.Cause(e)
	}
	if m[1] == "" {
		return m[2], nil
	}
	return m[1], nil
}
