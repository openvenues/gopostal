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

type DuplicateStatus int

func (d DuplicateStatus) String() string {
	switch d {
	case DuplicateStatusNull:
		return "Null Duplicate"
	case DuplicateStatusNon:
		return "Non Duplicate"
	case DuplicateStatusPossible:
		return "Possible Duplicate Needs Review"
	case DuplicateStatusLikely:
		return "Likely Duplicate"
	case DuplicateStatusExact:
		return "Exact Duplicate"
	}
	return "Unknown Duplicate Status"
}

type DuplicateOptions struct {
	Languages []string
}

type FuzzyDuplicateOptions struct {
	Languages            []string
	NeedsReviewThreshold float64
	LikelyDupeThreshold  float64
}

type FuzzyDuplicateStatus struct {
	DuplicateStatus DuplicateStatus
	Similarity      float64
}

func DefaultFuzzyDuplicateOptions() FuzzyDuplicateOptions {
	cDefaultFuzzyDuplicateOptions := C.libpostal_get_default_fuzzy_duplicate_options()
	return FuzzyDuplicateOptions{
		NeedsReviewThreshold: float64(cDefaultFuzzyDuplicateOptions.needs_review_threshold),
		LikelyDupeThreshold:  float64(cDefaultFuzzyDuplicateOptions.likely_dupe_threshold),
	}
}

const (
	DuplicateStatusNull     DuplicateStatus = C.LIBPOSTAL_NULL_DUPLICATE_STATUS
	DuplicateStatusNon      DuplicateStatus = C.LIBPOSTAL_NON_DUPLICATE
	DuplicateStatusPossible DuplicateStatus = C.LIBPOSTAL_POSSIBLE_DUPLICATE_NEEDS_REVIEW
	DuplicateStatusLikely   DuplicateStatus = C.LIBPOSTAL_LIKELY_DUPLICATE
	DuplicateStatusExact    DuplicateStatus = C.LIBPOSTAL_EXACT_DUPLICATE
)

func DefaultDuplicateOptions() DuplicateOptions {
	return DuplicateOptions{}
}

func IsDuplicate(addressComponent AddressComponent, value1, value2 string, options DuplicateOptions) DuplicateStatus {
	cValue1 := C.CString(value1)
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := C.CString(value2)
	defer C.free(unsafe.Pointer(cValue2))

	cOptions := C.libpostal_get_default_duplicate_options()

	if options.Languages != nil {
		cLanguages := make([]*C.char, len(options.Languages))

		for i := 0; i < len(options.Languages); i++ {
			cLang := C.CString(options.Languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}

		cOptions.languages = &cLanguages[0]
		cOptions.num_languages = C.size_t(len(options.Languages))
	} else {
		cOptions.num_languages = 0
	}

	switch addressComponent {
	case AddressStreet:
		return DuplicateStatus(C.libpostal_is_street_duplicate(cValue1, cValue2, cOptions))
	case AddressName:
		return DuplicateStatus(C.libpostal_is_name_duplicate(cValue1, cValue2, cOptions))
	case AddressHouseNumber:
		return DuplicateStatus(C.libpostal_is_house_number_duplicate(cValue1, cValue2, cOptions))
	case AddressPoBox:
		return DuplicateStatus(C.libpostal_is_po_box_duplicate(cValue1, cValue2, cOptions))
	case AddressUnit:
		return DuplicateStatus(C.libpostal_is_unit_duplicate(cValue1, cValue2, cOptions))
	case AddressPostalCode:
		return DuplicateStatus(C.libpostal_is_postal_code_duplicate(cValue1, cValue2, cOptions))
	}

	return DuplicateStatusNull
}

func IsNameDuplicateFuzzy(tokens1 []string, scores1 float64, tokens2 []string, scores2 float64, options FuzzyDuplicateOptions) FuzzyDuplicateStatus {
	return isDuplicateFuzzy(AddressName, tokens1, scores1, tokens2, scores2, options)
}

func IsStreetDuplicateFuzzy(tokens1 []string, scores1 float64, tokens2 []string, scores2 float64, options FuzzyDuplicateOptions) FuzzyDuplicateStatus {
	return isDuplicateFuzzy(AddressStreet, tokens1, scores1, tokens2, scores2, options)
}

func isDuplicateFuzzy(addressComponent AddressComponent, tokens1 []string, scores1 float64, tokens2 []string, scores2 float64, options FuzzyDuplicateOptions) FuzzyDuplicateStatus {
	cOptions := C.libpostal_get_default_fuzzy_duplicate_options()
	cOptions.needs_review_threshold = C.double(options.NeedsReviewThreshold)
	cOptions.likely_dupe_threshold = C.double(options.LikelyDupeThreshold)

	if options.Languages != nil {
		cLanguages := make([]*C.char, len(options.Languages))

		for i := 0; i < len(options.Languages); i++ {
			cLang := C.CString(options.Languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}

		cOptions.languages = &cLanguages[0]
		cOptions.num_languages = C.size_t(len(options.Languages))
	}

	cTokens1 := make([]*C.char, len(tokens1))
	for i, token1 := range tokens1 {
		cToken1 := C.CString(token1)
		defer C.free(unsafe.Pointer(cToken1))
		cTokens1[i] = cToken1
	}
	cTokens2 := make([]*C.char, len(tokens2))
	for i, token2 := range tokens2 {
		cToken2 := C.CString(token2)
		defer C.free(unsafe.Pointer(cToken2))
		cTokens2[i] = cToken2
	}

	cNumTokens1 := C.ulong(len(tokens1))
	cNumTokens2 := C.ulong(len(tokens2))

	cScores1 := C.double(scores1)
	cScores2 := C.double(scores2)

	var cStatus C.libpostal_fuzzy_duplicate_status_t
	switch addressComponent {
	case AddressStreet:
		cStatus = C.libpostal_is_street_duplicate_fuzzy(cNumTokens1, &cTokens1[0], &cScores1, cNumTokens2, &cTokens2[0], &cScores2, cOptions)
	case AddressName:
		cStatus = C.libpostal_is_name_duplicate_fuzzy(cNumTokens1, &cTokens1[0], &cScores1, cNumTokens2, &cTokens2[0], &cScores2, cOptions)
	default:
		return FuzzyDuplicateStatus{}
	}

	return FuzzyDuplicateStatus{
		DuplicateStatus: DuplicateStatus(cStatus.status),
		Similarity:      float64(cStatus.similarity),
	}
}

func IsToponymDuplicate(labels1 []string, values1 []string, labels2 []string, values2 []string, options DuplicateOptions) DuplicateStatus {
	cOptions := C.libpostal_get_default_duplicate_options()
	if options.Languages != nil {
		cLanguages := make([]*C.char, len(options.Languages))
		for i := 0; i < len(options.Languages); i++ {
			cLang := C.CString(options.Languages[i])
			defer C.free(unsafe.Pointer(cLang))
			cLanguages[i] = cLang
		}
		cOptions.languages = &cLanguages[0]
		cOptions.num_languages = C.size_t(len(options.Languages))
	} else {
		cOptions.num_languages = 0
	}

	cLabels1 := make([]*C.char, len(labels1))
	for i, label := range labels1 {
		cLabel := C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
		cLabels1[i] = cLabel
	}
	cLabels2 := make([]*C.char, len(labels2))
	for i, label := range labels2 {
		cLabel := C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
		cLabels2[i] = cLabel
	}

	cValues1 := make([]*C.char, len(values1))
	for i, value := range values1 {
		cValue := C.CString(value)
		defer C.free(unsafe.Pointer(cValue))
		cValues1[i] = cValue
	}
	cValues2 := make([]*C.char, len(values2))
	for i, value := range values2 {
		cValue := C.CString(value)
		defer C.free(unsafe.Pointer(cValue))
		cValues2[i] = cValue
	}

	cNumComponents1 := C.ulong(len(labels1))
	cNumComponents2 := C.ulong(len(labels2))

	cStatus := C.libpostal_is_toponym_duplicate(cNumComponents1, &cLabels1[0], &cValues1[0], cNumComponents2, &cLabels2[0], &cValues2[0], cOptions)

	return DuplicateStatus(cStatus)
}
