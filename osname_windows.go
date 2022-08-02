//go:build windows

package osname

import "golang.org/x/sys/windows/registry"

func osname() (string, error) {
	k, e := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if e != nil {
		return "", _ReadRegistryFailed.Cause(e)
	}
	s, _, e := k.GetStringValue("ProductName")
	if e != nil {
		return "", _ReadRegistryFailed.Cause(e)
	}
	return s, nil
}
