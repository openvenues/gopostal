package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"log"
)

func init() {
	err := Setup()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Setup() error {
	if !bool(C.libpostal_setup()) {
		return fmt.Errorf("Could not load libpostal_setup")
	}
	if !bool(C.libpostal_setup_language_classifier()) {
		return fmt.Errorf("Could not load libpostal_setup_language_classifier")
	}
	if !bool(C.libpostal_setup_parser()) {
		return fmt.Errorf("Could not load libpostal_setup_parser")
	}

	return nil
}

func Teardown() {
	C.libpostal_teardown()
	C.libpostal_teardown_language_classifier()
	C.libpostal_teardown_parser()
}

const (
	AddressLabelHouse         = "house"
	AddressLabelHouseNumber   = "house_number"
	AddressLabelPoBox         = "po_box"
	AddressLabelBuilding      = "building"
	AddressLabelEntrance      = "entrance"
	AddressLabelStaircase     = "staircase"
	AddressLabelLevel         = "level"
	AddressLabelUnit          = "unit"
	AddressLabelRoad          = "road"
	AddressLabelMetroStation  = "metro_station"
	AddressLabelSuburb        = "suburb"
	AddressLabelCityDistrict  = "city_district"
	AddressLabelCity          = "city"
	AddressLabelStateDistrict = "state_district"
	AddressLabelIsland        = "island"
	AddressLabelState         = "state"
	AddressLabelPostalCode    = "postcode"
	AddressLabelCountryRegion = "country_region"
	AddressLabelCountry       = "country"
	AddressLabelWorldRegion   = "world_region"
	AddressLabelWebsite       = "website"
	AddressLabelTelephone     = "phone"
)
