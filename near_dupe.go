package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type NearDupeHashOptions struct {
	WithName                      bool
	WithAddress                   bool
	WithUnit                      bool
	WithCityOrEquivalent          bool
	WithSmallContainingBoundaries bool
	WithPostalCode                bool
	WithLatLon                    bool
	Latitude                      float64
	Longitude                     float64
	GeohashPrecision              uint32
	NameAndAddressKeys            bool
	NameOnlyKeys                  bool
	AddressOnlyKeys               bool
	Languages                     []string
}

func DefaultNearDupeHashOptions() NearDupeHashOptions {
	cDefaultNearDupeHashOptions := C.libpostal_get_near_dupe_hash_default_options()
	return NearDupeHashOptions{
		WithName:                      bool(cDefaultNearDupeHashOptions.with_name),
		WithAddress:                   bool(cDefaultNearDupeHashOptions.with_address),
		WithUnit:                      bool(cDefaultNearDupeHashOptions.with_unit),
		WithCityOrEquivalent:          bool(cDefaultNearDupeHashOptions.with_city_or_equivalent),
		WithSmallContainingBoundaries: bool(cDefaultNearDupeHashOptions.with_small_containing_boundaries),
		WithPostalCode:                bool(cDefaultNearDupeHashOptions.with_postal_code),
		WithLatLon:                    bool(cDefaultNearDupeHashOptions.with_latlon),
		Latitude:                      float64(cDefaultNearDupeHashOptions.latitude),
		Longitude:                     float64(cDefaultNearDupeHashOptions.longitude),
		GeohashPrecision:              uint32(cDefaultNearDupeHashOptions.geohash_precision),
		NameAndAddressKeys:            bool(cDefaultNearDupeHashOptions.name_and_address_keys),
		NameOnlyKeys:                  bool(cDefaultNearDupeHashOptions.name_only_keys),
		AddressOnlyKeys:               bool(cDefaultNearDupeHashOptions.address_only_keys),
		Languages:                     nil,
	}
}

func NearDupeHashes(labels []string, values []string, options NearDupeHashOptions) []string {
	cOptions := C.libpostal_get_near_dupe_hash_default_options()
	cOptions.with_name = C.bool(options.WithName)
	cOptions.with_address = C.bool(options.WithAddress)
	cOptions.with_unit = C.bool(options.WithUnit)
	cOptions.with_city_or_equivalent = C.bool(options.WithCityOrEquivalent)
	cOptions.with_small_containing_boundaries = C.bool(options.WithSmallContainingBoundaries)
	cOptions.with_postal_code = C.bool(options.WithPostalCode)
	cOptions.with_latlon = C.bool(options.WithLatLon)
	cOptions.latitude = C.double(options.Latitude)
	cOptions.longitude = C.double(options.Longitude)
	cOptions.geohash_precision = C.uint32_t(options.GeohashPrecision)
	cOptions.name_and_address_keys = C.bool(options.NameAndAddressKeys)
	cOptions.name_only_keys = C.bool(options.NameOnlyKeys)
	cOptions.address_only_keys = C.bool(options.AddressOnlyKeys)

	cNumLanguages := C.size_t(0)
	var cLanguages []*C.char

	if options.Languages != nil {
		cLanguages = make([]*C.char, len(options.Languages))
		cNumLanguages = C.size_t(len(options.Languages))

		for i := 0; i < len(options.Languages); i++ {
			cLang := C.CString(options.Languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}
	}

	cLabels := make([]*C.char, len(labels))
	for i, label := range labels {
		cLabel := C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
		cLabels[i] = cLabel
	}

	cValues := make([]*C.char, len(values))
	for i, value := range values {
		cValue := C.CString(value)
		defer C.free(unsafe.Pointer(cValue))
		cValues[i] = cValue
	}

	cNumComponents := C.ulong(len(labels))
	cNumHashes := C.size_t(0)

	var cHashes **C.char
	if cNumLanguages > 0 {
		cHashes = C.libpostal_near_dupe_hashes_languages(cNumComponents, &cLabels[0], &cValues[0], cOptions, cNumLanguages, &cLanguages[0], &cNumHashes)
	} else {
		cHashes = C.libpostal_near_dupe_hashes(cNumComponents, &cLabels[0], &cValues[0], cOptions, &cNumHashes)
	}
	numHashes := uint64(cNumHashes)

	cHashesPtr := (*[1 << 28](*C.char))(unsafe.Pointer(cHashes))

	var hashes []string
	var i uint64
	for i = 0; i < numHashes; i++ {
		hashes = append(hashes, C.GoString(cHashesPtr[i]))
	}

	C.libpostal_expansion_array_destroy(cHashes, cNumHashes)

	return hashes
}
