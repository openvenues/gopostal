package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestNormalizedTokens(t *testing.T) {
	input := "ABC."
	token := postal.NormalizedTokens(input, postal.NormalizeStringOptionLowercase, postal.NormalizeTokenOptionDeleteFinalPeriod, true, nil)
	assert.Equal(t, "abc", token.String)
}

func TestNormalizeString(t *testing.T) {
	input := "ABC def"
	output := postal.NormalizeString(input, postal.NormalizeStringOptionLowercase, nil)
	assert.Equal(t, "abc def", output)
}
