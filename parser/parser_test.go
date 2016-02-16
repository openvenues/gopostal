package postal

import (
    "reflect"
    "testing"
)

func testParse(t *testing.T, address string, expectedOutput []ParsedComponent) {
    parsedComponents := ParseAddress(address)

    if len(parsedComponents) != len(expectedOutput) || !reflect.DeepEqual(parsedComponents, expectedOutput) {
        t.Error("parsed != expected", parsedComponents, "!=", expectedOutput)
    }
}

func TestParseUSAddress(t *testing.T) {
    t.Log("Testing US address")

    testParse(t, "781 Franklin Ave Crown Heights Brooklyn NYC NY 11216 USA", 
              []ParsedComponent {
                  {"781", "house_number"},
                  {"franklin ave", "road"},
                  {"crown heights", "suburb"},
                  {"brooklyn", "city_district"},
                  {"nyc", "city"},
                  {"ny", "state"},
                  {"11216", "postcode"},
                  {"usa", "country"},
              },
              )

    testParse(t, "whole foods ny",
              []ParsedComponent {
                {"whole foods", "house"},
                {"ny", "state"},
              },
              )

    testParse(t, "1917/2 Pike Drive",
              []ParsedComponent {
                {"1917 / 2", "house_number"},
                {"pike drive", "road"},
              },
              )


    testParse(t, "3437 warwickshire rd,pa",
              []ParsedComponent {
                {"3437", "house_number"},
                {"warwickshire rd", "road"},
                {"pa", "state"},
              },
              )

    testParse(t, "3437 warwickshire rd, pa",
              []ParsedComponent {
                {"3437", "house_number"},
                {"warwickshire rd", "road"},
                {"pa", "state"},
              },
              )

    testParse(t, "3437 warwickshire rd pa",
              []ParsedComponent {
                {"3437", "house_number"},
                {"warwickshire rd", "road"},
                {"pa", "state"},
              },
              )

}
