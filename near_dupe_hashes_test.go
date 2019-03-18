package postal_test

import (
	"fmt"
	"testing"

	"github.com/skyline-ai/postal"
)

func TestNearDupeHashes(t *testing.T) {
	lables := []string{postal.AddressParserLabelRoad, postal.AddressParserLabelHouseNumber, postal.AddressParserLabelCity, postal.AddressParserLabelState}
	values := []string{"east beaver creek rd", "426", "knoxville", "tn"}
	opts := postal.DefaultNearDupeHashOptions
	opts.AddressOnlyKeys = true
	opts.WithLatLon = true
	opts.Latitude = 35.85821
	opts.Longitude = -84.08088
	hashes := postal.NearDupeHashes(lables, values, opts)
	for _, hash := range hashes {
		fmt.Println(hash)
	}
}
