package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>

*/
import "C"

import (
    "unsafe"
    "log"
)

func init() {
    if (!bool(C.libpostal_setup()) || !bool(C.libpostal_setup_language_classifier())) {
        log.Fatal("Could not load libpostal")
    }
}

const (
    AddressAny = C.ADDRESS_ANY
    AddressName = C.ADDRESS_NAME
    AddressHouseNumber = C.ADDRESS_HOUSE_NUMBER
    AddressStreet = C.ADDRESS_STREET
    AddressUnit = C.ADDRESS_UNIT
    AddressLocality = C.ADDRESS_LOCALITY
    AddressAdmin1 = C.ADDRESS_ADMIN1
    AddressAdmin2 = C.ADDRESS_ADMIN2
    AddressAdmin3 = C.ADDRESS_ADMIN3
    AddressAdmin4 = C.ADDRESS_ADMIN4
    AddressAdminOther = C.ADDRESS_ADMIN_OTHER
    AddressCountry = C.ADDRESS_COUNTRY
    AddressNeighborhood = C.ADDRESS_NEIGHBORHOOD
    AddressAll = C.ADDRESS_ALL
)

type ExpandOptions struct {
    languages []string
    addressComponents uint16
    latinAscii bool
    transliterate bool
    stripAccents bool
    decompose bool
    lowercase bool
    trimString bool
    replaceWordHyphens bool
    deleteWordHyphens bool
    replaceNumericHyphens bool
    deleteNumericHyphens bool
    splitAlphaFromNumeric bool
    deleteFinalPeriods bool
    deleteAcronymPeriods bool
    dropEnglishPossessives bool
    deleteApostrophes bool
    expandNumex bool
    romanNumerals bool
}

var cDefaultOptions = C.get_libpostal_default_options()

func getDefaultExpansionOptions() ExpandOptions {
    return ExpandOptions{
        languages: nil,
        addressComponents: uint16(cDefaultOptions.address_components),
        latinAscii: bool(cDefaultOptions.latin_ascii),
        transliterate: bool(cDefaultOptions.transliterate),
        stripAccents: bool(cDefaultOptions.strip_accents),
        decompose: bool(cDefaultOptions.decompose),
        lowercase: bool(cDefaultOptions.lowercase),
        trimString: bool(cDefaultOptions.trim_string),
        replaceWordHyphens: bool(cDefaultOptions.replace_word_hyphens),
        deleteWordHyphens: bool(cDefaultOptions.delete_word_hyphens),
        replaceNumericHyphens: bool(cDefaultOptions.replace_numeric_hyphens),
        deleteNumericHyphens: bool(cDefaultOptions.delete_numeric_hyphens),
        splitAlphaFromNumeric: bool(cDefaultOptions.split_alpha_from_numeric),
        deleteFinalPeriods: bool(cDefaultOptions.delete_final_periods),
        deleteAcronymPeriods: bool(cDefaultOptions.delete_acronym_periods),
        dropEnglishPossessives: bool(cDefaultOptions.drop_english_possessives),
        deleteApostrophes: bool(cDefaultOptions.delete_apostrophes),
        expandNumex: bool(cDefaultOptions.expand_numex),
        romanNumerals: bool(cDefaultOptions.roman_numerals),
    }
}

var libpostalDefaultOptions = getDefaultExpansionOptions()

func ExpandAddressOptions(address string, options ExpandOptions) []string {
    cAddress := C.CString(address)
    defer C.free(unsafe.Pointer(cAddress))

    var char_ptr *C.char
    ptr_size := unsafe.Sizeof(char_ptr)

    cOptions := C.get_libpostal_default_options()
    if options.languages != nil {
        cLanguages := C.calloc(C.size_t(len(options.languages)), C.size_t(ptr_size))
        cLanguagesPtr := (*[1<<30](*C.char))(unsafe.Pointer(cLanguages))

        defer C.free(unsafe.Pointer(cLanguages))

        for i := 0; i < len(options.languages); i++ {
            cLang := C.CString(options.languages[i])
            defer C.free(unsafe.Pointer(cLang))
            cLanguagesPtr[i] = cLang
        }

        cOptions.languages = (**C.char)(cLanguages)
        cOptions.num_languages = C.int(len(options.languages))
    } else {
        cOptions.num_languages = 0
    }

    cOptions.address_components = C.uint16_t(options.addressComponents)
    cOptions.latin_ascii = C.bool(options.latinAscii)
    cOptions.transliterate = C.bool(options.transliterate)
    cOptions.strip_accents = C.bool(options.stripAccents)
    cOptions.decompose = C.bool(options.decompose)
    cOptions.lowercase = C.bool(options.lowercase)
    cOptions.trim_string = C.bool(options.trimString)
    cOptions.replace_word_hyphens = C.bool(options.replaceWordHyphens)
    cOptions.delete_word_hyphens = C.bool(options.deleteWordHyphens)
    cOptions.replace_numeric_hyphens = C.bool(options.replaceNumericHyphens)
    cOptions.delete_numeric_hyphens = C.bool(options.deleteNumericHyphens)
    cOptions.split_alpha_from_numeric = C.bool(options.splitAlphaFromNumeric)
    cOptions.delete_final_periods = C.bool(options.deleteFinalPeriods)
    cOptions.delete_acronym_periods = C.bool(options.deleteAcronymPeriods)
    cOptions.drop_english_possessives = C.bool(options.dropEnglishPossessives)
    cOptions.delete_apostrophes = C.bool(options.deleteApostrophes)
    cOptions.expand_numex = C.bool(options.expandNumex)
    cOptions.roman_numerals = C.bool(options.romanNumerals)

    var cNumExpansions = C.size_t(0)

    cExpansions := C.expand_address(cAddress, cOptions, &cNumExpansions)

    numExpansions := uint64(cNumExpansions)

    var expansions = make([]string, numExpansions)

    // Accessing a C array
    cExpansionsPtr := (*[1<<30](*C.char))(unsafe.Pointer(cExpansions))

    var i uint64
    for i = 0; i < numExpansions; i++ {
        expansions[i] = C.GoString(cExpansionsPtr[i])
    }

    return expansions
}

func ExpandAddress(address string) []string {
    expansions := ExpandAddressOptions(address, libpostalDefaultOptions);
    return expansions
}

