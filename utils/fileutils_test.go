package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTempDir(t *testing.T) {
	fd, err := CreateTempFile()
	assert.NoError(t, err)
	assert.NoError(t, fd.Close())
	err = os.Remove(fd.Name())
	assert.NoError(t, err)
}
