package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSysInfo(t *testing.T) {
	ast := assert.New(t)
	sysInfo := NewSysInfo()
	ast.NotNil(sysInfo)
}
