package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
	"github.com/stretchr/testify/assert"
)

func TestPlaceLanguages(t *testing.T) {
	lables := []string{postal.AddressLabelRoad, postal.AddressLabelHouseNumber, postal.AddressLabelCity, postal.AddressLabelState}
	values := []string{"east beaver creek rd", "426", "knoxville", "tn"}
	languages := postal.PlaceLanguages(lables, values)
	assert.Equal(t, []string{"en"}, languages)
}
