package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	assert.Implements(t, (*Logger)(nil), &zapSugaredLoggerWrapper{})
}
