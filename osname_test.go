package osname

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOsName(t *testing.T) {
	s, e := OsName()
	assert.NoError(t, e)
	fmt.Println(s)
	t.Log(s)
}
