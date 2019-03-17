package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	tokens := postal.Tokenize("426 E Beaver Creek Dr, Knoxville, TN 37918, USA", false)
	assert.Len(t, tokens, 12)
}
