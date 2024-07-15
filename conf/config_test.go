package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureApiVersion(t *testing.T) {
	Config.Host = "https://disco.test/api/public/v1/"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/api/public/v1", Config.Host)
	Config.Host = "https://disco.test/api/public/v1"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/api/public/v1", Config.Host)
	Config.Host = "https://disco.test/api/public/"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/api/public/v1", Config.Host)
	Config.Host = "https://disco.test/api/public/v32"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/api/public/v32", Config.Host)
	Config.Host = "https://disco.test/disco"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/disco/v1", Config.Host)
	Config.Host = "https://disco.test/disco/v1"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/disco/v1", Config.Host)
	Config.Host = "https://disco.test/disco/"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/disco/v1", Config.Host)
	Config.Host = "https://disco.test/disco/v32"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/disco/v32", Config.Host)
}
