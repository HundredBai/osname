package osname

const ErrGetOsName = err("ErrGetOsName")
const (
	_ReadPlistFailed     = err("Read plist failed")
	_ReadOsReleaseFailed = err("Read os-release failed")
	_ReadRegistryFailed  = err("Read registry failed")
)

func OsName() (string, error) {
	s, e := osname()
	if e != nil {
		return "", ErrGetOsName.Cause(e)
	}
	return s, nil
}
