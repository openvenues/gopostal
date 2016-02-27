package postal

import (
	"encoding/json"
    "reflect"
    "testing"
)

func testParse(t *testing.T, address string, expectedOutput []ParsedComponent, expectedJSON string) {
    parsedComponents := ParseAddress(address)

    if len(parsedComponents) != len(expectedOutput) || !reflect.DeepEqual(parsedComponents, expectedOutput) {
        t.Error("parsed != expected: ", parsedComponents, "!=", expectedOutput)
    }

	// Test JSON marshaling.
	marshaledJSON, err := json.Marshal(parsedComponents)
	if err != nil {
		t.Error("JSON.marshal error: " + err.Error())
	}

	if string(marshaledJSON) != expectedJSON {
        t.Error("json != expected: ", string(marshaledJSON), "!=", expectedJSON)
	}

	// Test JSON unmarshaling.
	var unmarshaledComponents []ParsedComponent
	if err := json.Unmarshal(marshaledJSON, &unmarshaledComponents); err != nil {
		t.Error("JSON.unmarshal error: " + err.Error())
	}
    if !reflect.DeepEqual(unmarshaledComponents, expectedOutput) {
        t.Error("unmarshaled != expected: ", unmarshaledComponents, "!=", expectedOutput)
    }
}

func TestParseUSAddress(t *testing.T) {
    t.Log("Testing US address")

    testParse(t, "781 Franklin Ave Crown Heights Brooklyn NYC NY 11216 USA", 
              []ParsedComponent {
                  {"house_number", "781"},
                  {"road", "franklin ave"},
                  {"suburb", "crown heights"},
                  {"city_district", "brooklyn"},
                  {"city", "nyc"},
                  {"state", "ny"},
                  {"postcode", "11216"},
                  {"country", "usa"},
              },
              `[{"label":"house_number","value":"781"},{"label":"road","value":"franklin ave"},{"label":"suburb","value":"crown heights"},{"label":"city_district","value":"brooklyn"},{"label":"city","value":"nyc"},{"label":"state","value":"ny"},{"label":"postcode","value":"11216"},{"label":"country","value":"usa"}]`,
              )

    testParse(t, "whole foods ny",
              []ParsedComponent {
                {"house", "whole foods"},
                {"state", "ny"},
              },
              `[{"label":"house","value":"whole foods"},{"label":"state","value":"ny"}]`,
              )

    testParse(t, "1917/2 Pike Drive",
              []ParsedComponent {
                {"house_number", "1917 / 2"},
                {"road", "pike drive"},
              },
              `[{"label":"house_number","value":"1917 / 2"},{"label":"road","value":"pike drive"}]`,
              )

    testParse(t, "3437 warwickshire rd,pa",
              []ParsedComponent {
                {"house_number", "3437"},
                {"road", "warwickshire rd"},
                {"state", "pa"},
              },
              `[{"label":"house_number","value":"3437"},{"label":"road","value":"warwickshire rd"},{"label":"state","value":"pa"}]`,
              )

    testParse(t, "3437 warwickshire rd, pa",
              []ParsedComponent {
                {"house_number", "3437"},
                {"road", "warwickshire rd"},
                {"state", "pa"},
              },
              `[{"label":"house_number","value":"3437"},{"label":"road","value":"warwickshire rd"},{"label":"state","value":"pa"}]`,
              )

    testParse(t, "3437 warwickshire rd pa",
              []ParsedComponent {
                {"house_number", "3437"},
                {"road", "warwickshire rd"},
                {"state", "pa"},
              },
              `[{"label":"house_number","value":"3437"},{"label":"road","value":"warwickshire rd"},{"label":"state","value":"pa"}]`,
              )
}
