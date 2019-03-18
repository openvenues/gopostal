package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestNearDupeHashes(t *testing.T) {
	lables := []string{postal.AddressLabelRoad, postal.AddressLabelHouseNumber, postal.AddressLabelCity, postal.AddressLabelState}
	values := []string{"east beaver creek rd", "426", "knoxville", "tn"}
	opts := postal.DefaultNearDupeHashOptions()
	opts.AddressOnlyKeys = true
	opts.WithLatLon = true
	opts.Latitude = 35.85821
	opts.Longitude = -84.08088
	opts.Languages = []string{"en"}
	hashes := postal.NearDupeHashes(lables, values, opts)
	assert.Len(t, hashes, 20)
}
