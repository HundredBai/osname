//go:build linux

package osname

import (
	"io/ioutil"
	"regexp"
)

func osname() (string, error) {
	data, e := ioutil.ReadFile("/etc/os-release")
	if e == nil {
		m := regexp.MustCompile("PRETTY_NAME\\s*=\\s*(?:\\\"([^\\\"\\n]+)\\\"|([^\\n]+))").FindStringSubmatch(string(data))
		if m != nil {
			if m[1] == "" {
				return m[2], nil
			}
			return m[1], nil
		}
	}
	// 兼容 centos6: /etc/os-release 不存在时，尝试读取 /etc/redhat-release
	data, e = ioutil.ReadFile("/etc/redhat-release")
	if e == nil {
		return string(data), nil
	}
	return "", _ReadOsReleaseFailed.Cause(e)
}
