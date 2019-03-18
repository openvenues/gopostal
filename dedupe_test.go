package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestIsDuplicate(t *testing.T) {
	status := postal.IsDuplicate(postal.AddressStreet, "6020 Churchland St #1", "6020 Churchland Blvd #1", postal.DefaultDuplicateOptions())
	assert.Equal(t, postal.DuplicateStatusPossible, status)
	status = postal.IsDuplicate(postal.AddressName, "Home I", "Home 2", postal.DefaultDuplicateOptions())
	assert.Equal(t, postal.DuplicateStatusNon, status)
}

func TestIsToponymDuplicate(t *testing.T) {
	labels1 := []string{postal.AddressLabelRoad, postal.AddressLabelCity, postal.AddressLabelState, postal.AddressLabelCountry, postal.AddressLabelHouse}
	values1 := []string{"426 East Beaver Creek Rd", "Knoxville", "TN", "USA", "home 1"}
	labels2 := []string{postal.AddressLabelRoad, postal.AddressLabelCity, postal.AddressLabelState, postal.AddressLabelCountry, postal.AddressLabelHouse}
	values2 := []string{"426 East Beaver Creek Rd", "Knoxville", "TN", "USA", "home1"}
	status := postal.IsToponymDuplicate(labels1, values1, labels2, values2, postal.DefaultDuplicateOptions())
	assert.Equal(t, postal.DuplicateStatusExact, status)
}

func TestIsNameDuplicateFuzzy(t *testing.T) {
	tokens1 := []string{"The", "Name 1"}
	scores1 := 1.0
	tokens2 := []string{"The", "Name 2"}
	scores2 := 1.0
	opts := postal.DefaultFuzzyDuplicateOptions()
	status := postal.IsNameDuplicateFuzzy(tokens1, scores1, tokens2, scores2, opts)
	assert.Equal(t, postal.DuplicateStatusLikely.String(), status.DuplicateStatus.String())
}

func TestIsStreetDuplicateFuzzy(t *testing.T) {
	tokens1 := []string{"East", "Beaver", "Creek", "Rd"}
	scores1 := 1.0
	tokens2 := []string{"East", "Beaver", "Creek", "Road"}
	scores2 := 1.0
	opts := postal.DefaultFuzzyDuplicateOptions()
	status := postal.IsStreetDuplicateFuzzy(tokens1, scores1, tokens2, scores2, opts)
	assert.Equal(t, postal.DuplicateStatusPossible.String(), status.DuplicateStatus.String())
}
