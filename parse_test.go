package phpregexp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	_, err := Compile("")
	assert.Equal(ErrMalformedRegexp, err)
	_, err = Compile(" ")
	assert.Equal(ErrMalformedRegexp, err)
	_, err = Compile("  ")
	assert.Equal(ErrMalformedRegexp, err)

	_, err = Compile("/[a-z]+")
	assert.Equal(ErrCantFindClosingSeparator, err)
	_, err = Compile("/[a-z]//+")
	assert.Equal(ErrCantFindClosingSeparator, err)

	r, err := Compile("/[a-z]+/")
	assert.NotNil(r)
	assert.NoError(err)
	assert.Equal("[a-z]+", r.String())
	r, err = Compile("/[a-z]+/i")
	assert.NotNil(r)
	assert.NoError(err)
	assert.Equal("(?i)[a-z]+", r.String())
	r, err = Compile("/[a-z]+/m")
	assert.NotNil(r)
	assert.NoError(err)
	assert.Equal("(?m)[a-z]+", r.String())
	r, err = Compile("/[a-z]+/sU")
	assert.NotNil(r)
	assert.NoError(err)
	assert.Equal("(?sU)[a-z]+", r.String())
	r, err = Compile("[a-]+")
	assert.Nil(r)
	assert.Error(err)
}
