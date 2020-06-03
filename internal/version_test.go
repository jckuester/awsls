package internal_test

import (
	"runtime"
	"testing"

	"github.com/jckuester/awsls/internal"
	"github.com/stretchr/testify/assert"
)

func TestBuildVersionString(t *testing.T) {
	actualVersionString := internal.BuildVersionString()

	assert.Equal(t, actualVersionString, "version: dev\ncommit: ?\nbuilt at: ?\nusing: "+runtime.Version())
}
