package gigachat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientConnection(t *testing.T) {
	client := GetClient()
	assert.NotNil(t, client)
}