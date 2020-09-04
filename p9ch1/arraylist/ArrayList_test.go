package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayList_String(t *testing.T) {
	al := ArrayList{}
	al.Add(3)
	al.Add(5)

	assert.Equal(t, "[3 5]", al.String())
}
