package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"

type NormalizeOptions struct {
	Languages              []string
	AddressComponents      uint16
	LatinASCII             bool
	Transliterate          bool
	StripAccents           bool
	Decompose              bool
	Lowercase              bool
	TrimString             bool
	ReplaceWordHyphens     bool
	DeleteWordHyphens      bool
	ReplaceNumericHyphens  bool
	DeleteNumericHyphens   bool
	SplitAlphaFromNumeric  bool
	DeleteFinalPeriods     bool
	DeleteAcronymPeriods   bool
	DropEnglishPossessives bool
	DeleteApostrophes      bool
	ExpandNumex            bool
	RomanNumerals          bool
}

var (
	cDefaultNormalizeOptions = C.libpostal_get_default_options()
)

var DefaultNormalizeOptions = NormalizeOptions{
	Languages:              nil,
	AddressComponents:      uint16(cDefaultNormalizeOptions.address_components),
	LatinASCII:             bool(cDefaultNormalizeOptions.latin_ascii),
	Transliterate:          bool(cDefaultNormalizeOptions.transliterate),
	StripAccents:           bool(cDefaultNormalizeOptions.strip_accents),
	Decompose:              bool(cDefaultNormalizeOptions.decompose),
	Lowercase:              bool(cDefaultNormalizeOptions.lowercase),
	TrimString:             bool(cDefaultNormalizeOptions.trim_string),
	ReplaceWordHyphens:     bool(cDefaultNormalizeOptions.replace_word_hyphens),
	DeleteWordHyphens:      bool(cDefaultNormalizeOptions.delete_word_hyphens),
	ReplaceNumericHyphens:  bool(cDefaultNormalizeOptions.replace_numeric_hyphens),
	DeleteNumericHyphens:   bool(cDefaultNormalizeOptions.delete_numeric_hyphens),
	SplitAlphaFromNumeric:  bool(cDefaultNormalizeOptions.split_alpha_from_numeric),
	DeleteFinalPeriods:     bool(cDefaultNormalizeOptions.delete_final_periods),
	DeleteAcronymPeriods:   bool(cDefaultNormalizeOptions.delete_acronym_periods),
	DropEnglishPossessives: bool(cDefaultNormalizeOptions.drop_english_possessives),
	DeleteApostrophes:      bool(cDefaultNormalizeOptions.delete_apostrophes),
	ExpandNumex:            bool(cDefaultNormalizeOptions.expand_numex),
	RomanNumerals:          bool(cDefaultNormalizeOptions.roman_numerals),
}
