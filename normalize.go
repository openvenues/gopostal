package postal

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

var DefaultNormalizeOptions = NormalizeOptions{
	Languages:              nil,
	AddressComponents:      uint16(cDefaultOptions.address_components),
	LatinASCII:             bool(cDefaultOptions.latin_ascii),
	Transliterate:          bool(cDefaultOptions.transliterate),
	StripAccents:           bool(cDefaultOptions.strip_accents),
	Decompose:              bool(cDefaultOptions.decompose),
	Lowercase:              bool(cDefaultOptions.lowercase),
	TrimString:             bool(cDefaultOptions.trim_string),
	ReplaceWordHyphens:     bool(cDefaultOptions.replace_word_hyphens),
	DeleteWordHyphens:      bool(cDefaultOptions.delete_word_hyphens),
	ReplaceNumericHyphens:  bool(cDefaultOptions.replace_numeric_hyphens),
	DeleteNumericHyphens:   bool(cDefaultOptions.delete_numeric_hyphens),
	SplitAlphaFromNumeric:  bool(cDefaultOptions.split_alpha_from_numeric),
	DeleteFinalPeriods:     bool(cDefaultOptions.delete_final_periods),
	DeleteAcronymPeriods:   bool(cDefaultOptions.delete_acronym_periods),
	DropEnglishPossessives: bool(cDefaultOptions.drop_english_possessives),
	DeleteApostrophes:      bool(cDefaultOptions.delete_apostrophes),
	ExpandNumex:            bool(cDefaultOptions.expand_numex),
	RomanNumerals:          bool(cDefaultOptions.roman_numerals),
}
