package postal_test

import (
	"testing"

	"github.com/skyline-ai/postal"
)

func testExpansionInOutput(t *testing.T, address string, output string, expansions []string) {
	for i := 0; i < len(expansions); i++ {
		if expansions[i] == output {
			return
		}
	}

	t.Error("expansion", output, "not found in expansions for address", address)
}

func testExpansion(t *testing.T, address string, output string) {
	expansions := postal.ExpandAddress(address, postal.DefaultNormalizeOptions())
	testExpansionInOutput(t, address, output, expansions)
}

func testExpansionWithOptions(t *testing.T, address string, output string, options postal.NormalizeOptions) {
	expansions := postal.ExpandAddress(address, options)

	testExpansionInOutput(t, address, output, expansions)
}

func TestEnglishExpansions(t *testing.T) {
	testExpansion(t, "123 Main St", "123 main street")

	englishOptions := postal.DefaultNormalizeOptions()
	englishOptions.Languages = []string{"en"}

	testExpansionWithOptions(t, "30 West Twenty-sixth St Fl No. 7", "30 west 26th street floor number 7", englishOptions)
	testExpansionWithOptions(t, "Thirty W 26th St Fl #7", "30 west 26th street floor number 7", englishOptions)

}

func TestMultilingualExpansions(t *testing.T) {
	multilingualOptions := postal.DefaultNormalizeOptions()
	multilingualOptions.Languages = []string{"en", "fr", "de"}

	testExpansionWithOptions(t, "st", "sankt", multilingualOptions)
	testExpansionWithOptions(t, "st", "saint", multilingualOptions)
}

func TestNonASCIIExpansions(t *testing.T) {
	testExpansion(t, "Friedrichstraße 128, Berlin, Germany", "friedrich strasse 128 berlin germany")
}
