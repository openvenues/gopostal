package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"

import (
	"unicode/utf8"
	"unsafe"
)

func ExpandAddress(address string, options NormalizeOptions) []string {
	return expandAddress(address, false, options)
}

func ExpandAddressRoot(address string, options NormalizeOptions) []string {
	return expandAddress(address, true, options)
}

func expandAddress(address string, root bool, options NormalizeOptions) []string {
	if !utf8.ValidString(address) {
		return nil
	}

	cAddress := C.CString(address)
	defer C.free(unsafe.Pointer(cAddress))

	var charPtr *C.char
	ptrSize := unsafe.Sizeof(charPtr)

	cOptions := C.libpostal_get_default_options()
	if options.Languages != nil {
		cLanguages := C.calloc(C.size_t(len(options.Languages)), C.size_t(ptrSize))
		cLanguagesPtr := (*[1 << 30](*C.char))(unsafe.Pointer(cLanguages))

		defer C.free(unsafe.Pointer(cLanguages))

		for i := 0; i < len(options.Languages); i++ {
			cLang := C.CString(options.Languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguagesPtr[i] = cLang
		}

		cOptions.languages = (**C.char)(cLanguages)
		cOptions.num_languages = C.size_t(len(options.Languages))
	} else {
		cOptions.num_languages = 0
	}

	cOptions.address_components = C.uint16_t(options.AddressComponents)
	cOptions.latin_ascii = C.bool(options.LatinASCII)
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

	var cExpansions **C.char
	if root {
		cExpansions = C.libpostal_expand_address_root(cAddress, cOptions, &cNumExpansions)
	} else {
		cExpansions = C.libpostal_expand_address(cAddress, cOptions, &cNumExpansions)
	}

	numExpansions := uint64(cNumExpansions)

	var expansions = make([]string, numExpansions)

	cExpansionsPtr := (*[1 << 30](*C.char))(unsafe.Pointer(cExpansions))

	var i uint64
	for i = 0; i < numExpansions; i++ {
		expansions[i] = C.GoString(cExpansionsPtr[i])
	}

	C.libpostal_expansion_array_destroy(cExpansions, cNumExpansions)
	return expansions
}
