package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"

type AddressComponent uint16

const (
	AddressNone        AddressComponent = C.LIBPOSTAL_ADDRESS_NONE
	AddressAny         AddressComponent = C.LIBPOSTAL_ADDRESS_ANY
	AddressName        AddressComponent = C.LIBPOSTAL_ADDRESS_NAME
	AddressHouseNumber AddressComponent = C.LIBPOSTAL_ADDRESS_HOUSE_NUMBER
	AddressStreet      AddressComponent = C.LIBPOSTAL_ADDRESS_STREET
	AddressUnit        AddressComponent = C.LIBPOSTAL_ADDRESS_UNIT
	AddressLevel       AddressComponent = C.LIBPOSTAL_ADDRESS_LEVEL
	AddressStaircase   AddressComponent = C.LIBPOSTAL_ADDRESS_STAIRCASE
	AddressEntrance    AddressComponent = C.LIBPOSTAL_ADDRESS_ENTRANCE
	AddressCategory    AddressComponent = C.LIBPOSTAL_ADDRESS_CATEGORY
	AddressNear        AddressComponent = C.LIBPOSTAL_ADDRESS_NEAR
	AddressToponym     AddressComponent = C.LIBPOSTAL_ADDRESS_TOPONYM
	AddressPostalCode  AddressComponent = C.LIBPOSTAL_ADDRESS_POSTAL_CODE
	AddressPoBox       AddressComponent = C.LIBPOSTAL_ADDRESS_PO_BOX
	AddressAll         AddressComponent = C.LIBPOSTAL_ADDRESS_ALL
)
