package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureApiVersion(t *testing.T) {
	// Test that URLs without /public/v1 work correctly
	Config.Host = "https://disco.test/api/public/"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/api/public/v1", Config.Host)

	Config.Host = "https://disco.test/disco"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/disco/v1", Config.Host)

	Config.Host = "https://disco.test/disco/"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/disco/v1", Config.Host)

	Config.Host = "https://disco.test/disco/v32"
	EnsureApiVerison()
	assert.Equal(t, "https://disco.test/disco/v32", Config.Host)
}

// Note: Tests for URLs containing /public/v1 or /public/v2 will cause os.Exit(1)
// and cannot be tested in unit tests without mocking os.Exit
