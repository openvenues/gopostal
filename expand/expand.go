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
    "unicode/utf8"
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
    Languages []string
    AddressComponents uint16
    LatinAscii bool
    Transliterate bool
    StripAccents bool
    Decompose bool
    Lowercase bool
    TrimString bool
    ReplaceWordHyphens bool
    DeleteWordHyphens bool
    ReplaceNumericHyphens bool
    DeleteNumericHyphens bool
    SplitAlphaFromNumeric bool
    DeleteFinalPeriods bool
    DeleteAcronymPeriods bool
    DropEnglishPossessives bool
    DeleteApostrophes bool
    ExpandNumex bool
    RomanNumerals bool
}

var cDefaultOptions = C.get_libpostal_default_options()

func getDefaultExpansionOptions() ExpandOptions {
    return ExpandOptions{
        Languages: nil,
        AddressComponents: uint16(cDefaultOptions.address_components),
        LatinAscii: bool(cDefaultOptions.latin_ascii),
        Transliterate: bool(cDefaultOptions.transliterate),
        StripAccents: bool(cDefaultOptions.strip_accents),
        Decompose: bool(cDefaultOptions.decompose),
        Lowercase: bool(cDefaultOptions.lowercase),
        TrimString: bool(cDefaultOptions.trim_string),
        ReplaceWordHyphens: bool(cDefaultOptions.replace_word_hyphens),
        DeleteWordHyphens: bool(cDefaultOptions.delete_word_hyphens),
        ReplaceNumericHyphens: bool(cDefaultOptions.replace_numeric_hyphens),
        DeleteNumericHyphens: bool(cDefaultOptions.delete_numeric_hyphens),
        SplitAlphaFromNumeric: bool(cDefaultOptions.split_alpha_from_numeric),
        DeleteFinalPeriods: bool(cDefaultOptions.delete_final_periods),
        DeleteAcronymPeriods: bool(cDefaultOptions.delete_acronym_periods),
        DropEnglishPossessives: bool(cDefaultOptions.drop_english_possessives),
        DeleteApostrophes: bool(cDefaultOptions.delete_apostrophes),
        ExpandNumex: bool(cDefaultOptions.expand_numex),
        RomanNumerals: bool(cDefaultOptions.roman_numerals),
    }
}

var libpostalDefaultOptions = getDefaultExpansionOptions()

func ExpandAddressOptions(address string, options ExpandOptions) []string {
    if !utf8.ValidString(address) {
        return nil
    }

    cAddress := C.CString(address)
    defer C.free(unsafe.Pointer(cAddress))

    var char_ptr *C.char
    ptr_size := unsafe.Sizeof(char_ptr)

    cOptions := C.get_libpostal_default_options()
    if options.Languages != nil {
        cLanguages := C.calloc(C.size_t(len(options.Languages)), C.size_t(ptr_size))
        cLanguagesPtr := (*[1<<30](*C.char))(unsafe.Pointer(cLanguages))

        defer C.free(unsafe.Pointer(cLanguages))

        for i := 0; i < len(options.Languages); i++ {
            cLang := C.CString(options.Languages[i])
            defer C.free(unsafe.Pointer(cLang))
            cLanguagesPtr[i] = cLang
        }

        cOptions.languages = (**C.char)(cLanguages)
        cOptions.num_languages = C.int(len(options.Languages))
    } else {
        cOptions.num_languages = 0
    }

    cOptions.address_components = C.uint16_t(options.AddressComponents)
    cOptions.latin_ascii = C.bool(options.LatinAscii)
    cOptions.transliterate = C.bool(options.Transliterate)
    cOptions.strip_accents = C.bool(options.StripAccents)
    cOptions.decompose = C.bool(options.Decompose)
    cOptions.lowercase = C.bool(options.Lowercase)
    cOptions.trim_string = C.bool(options.TrimString)
    cOptions.replace_word_hyphens = C.bool(options.ReplaceWordHyphens)
    cOptions.delete_word_hyphens = C.bool(options.DeleteWordHyphens)
    cOptions.replace_numeric_hyphens = C.bool(options.ReplaceNumericHyphens)
    cOptions.delete_numeric_hyphens = C.bool(options.DeleteNumericHyphens)
    cOptions.split_alpha_from_numeric = C.bool(options.SplitAlphaFromNumeric)
    cOptions.delete_final_periods = C.bool(options.DeleteFinalPeriods)
    cOptions.delete_acronym_periods = C.bool(options.DeleteAcronymPeriods)
    cOptions.drop_english_possessives = C.bool(options.DropEnglishPossessives)
    cOptions.delete_apostrophes = C.bool(options.DeleteApostrophes)
    cOptions.expand_numex = C.bool(options.ExpandNumex)
    cOptions.roman_numerals = C.bool(options.RomanNumerals)

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
    return ExpandAddressOptions(address, libpostalDefaultOptions)
}

