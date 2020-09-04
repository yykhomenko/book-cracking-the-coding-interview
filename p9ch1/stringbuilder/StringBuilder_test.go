package stringbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringBuilder_Append(t *testing.T) {
	sb := StringBuilder{}
	sb.Append("Hello")
	sb.Append(" ")
	sb.Append("world")

	assert.Equal(t, "Hello world", sb.String())
}
