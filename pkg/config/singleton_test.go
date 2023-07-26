package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	singleton := GetInstance()

	assert.Equal(t, singleton.GetRootPath(), "")

	singleton.SetRootPath("/host")
	assert.Equal(t, singleton.GetRootPath(), "/host")

	path := singleton.AdjustPath("/dev/vhost-vdpa-0")
	assert.Equal(t, path, "/host/dev/vhost-vdpa-0")
}
