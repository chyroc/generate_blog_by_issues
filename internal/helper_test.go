package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64(t *testing.T) {
	assert.Equal(t, "IyDmoIfpopgKCi0gdGltZSAyMDE4LTAxLTI1", encodeBase64("# 标题\n\n- time 2018-01-25"))

	s, err := decodeBase64("IyBkZWZlcuS4jnJldHVybueahOmXrumimO+8iGRlZmVy5LmL5LiA77yJCgpb\nZ28taW50ZXJuYWxzLzAzLjQubWQgYXQgbWFzdGVyIMK3IHRpYW5jYWlhbWFv\n")
	assert.Nil(t, err)
	assert.Equal(t, "# defer与return的问题（defer之一）\n\n[go-internals/03.4.md at master · tiancaiamao", string(s))

	d, err := decodeBase64(encodeBase64("s"))
	assert.Nil(t, err)
	assert.Equal(t, "s", string(d))
}
