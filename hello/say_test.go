package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSay(t *testing.T) {
	assert.Equal(t, "saying: hi", Say("hi"))
}
