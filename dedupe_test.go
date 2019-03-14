package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestIsDuplicate(t *testing.T) {
	status := postal.IsDuplicate(postal.AddressStreet, "6020 Churchland St #1", "6020 Churchland Blvd #1", postal.GetDefaultDuplicateOptions())
	assert.Equal(t, postal.PossibleDuplicateNeedsReview, status)
	status = postal.IsDuplicate(postal.AddressName, "Home I", "Home 2", postal.GetDefaultDuplicateOptions())
	assert.Equal(t, postal.NonDuplicate, status)
}
