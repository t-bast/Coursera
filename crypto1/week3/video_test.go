package week3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t-bast/coursera/crypto1/cryptoutil"
	"github.com/t-bast/coursera/crypto1/week3"
)

func TestVideo(t *testing.T) {
	f, err := week3.NewFile("/path/to/a/file")
	assert.NoError(t, err, "week3.NewFile()")

	h := f.Hash()
	c := cryptoutil.Cipher(h)
	assert.Equal(t, "03c08f4ee0b576fe319338139c045c89c3e8e9409633bea29442e21425006ea8", c.String())
}
