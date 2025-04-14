// v0.1.0
// Author: wunderbarb

package imgpath

import (
	"testing"

	"github.com/wunderbarb/test"
)

func TestC3(t *testing.T) {
	_, assert := test.Describe(t)

	c := initC3("test6.png")

	for i := 0; i < len(c.path); i++ {
		p := c.Next()
		assert.Equal(0x10*uint8(i), p, "sample %d", i)
	}
	assert.Equal(c.Len(), 12)
}

func TestC4(t *testing.T) {
	_, assert := test.Describe(t)

	c := initC4("test7.png")
	assert.Equal(c.Len(), 16)
	for i := 0; i < len(c.path); i++ {
		p := c.Next()
		assert.Equal(0x10*uint8(i), p, "sample %d", i)
	}

}

func TestC5(t *testing.T) {
	_, assert := test.Describe(t)

	c := initC5("test8.png")
	assert.Equal(c.Len(), 20)
	for i := 0; i < len(c.path); i++ {
		p := c.Next()
		assert.Equal(0x10*uint8(i), p, "sample %d", i)
	}

}
