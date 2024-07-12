package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureApiVersion(t *testing.T) {
	Config.Host = "https://disco-int.app.corpintra.net/api/public/v1/"
	EnsureApiVerison()
	assert.Equal(t, "https://disco-int.app.corpintra.net/api/public/v1", Config.Host)
	Config.Host = "https://disco-int.app.corpintra.net/api/public/v1"
	EnsureApiVerison()
	assert.Equal(t, "https://disco-int.app.corpintra.net/api/public/v1", Config.Host)
	Config.Host = "https://disco-int.app.corpintra.net/api/public/"
	EnsureApiVerison()
	assert.Equal(t, "https://disco-int.app.corpintra.net/api/public/v1", Config.Host)
	Config.Host = "https://disco-int.app.corpintra.net/api/public/v32"
	EnsureApiVerison()
	assert.Equal(t, "https://disco-int.app.corpintra.net/api/public/v32", Config.Host)
}
