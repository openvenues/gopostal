package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type DuplicateStatus int
type DuplicateOptions struct {
	NumLanguages int
	Languages    []string
}

const (
	NullDuplicateStatus          DuplicateStatus = C.LIBPOSTAL_NULL_DUPLICATE_STATUS
	NonDuplicate                 DuplicateStatus = C.LIBPOSTAL_NON_DUPLICATE
	PossibleDuplicateNeedsReview DuplicateStatus = C.LIBPOSTAL_POSSIBLE_DUPLICATE_NEEDS_REVIEW
	LikelyDuplicate              DuplicateStatus = C.LIBPOSTAL_LIKELY_DUPLICATE
	ExactDuplicate               DuplicateStatus = C.LIBPOSTAL_EXACT_DUPLICATE
)

func GetDefaultDuplicateOptions() DuplicateOptions {
	return DuplicateOptions{
		NumLanguages: 0,
		Languages:    nil,
	}
}

func IsDuplicate(addressComponent uint16, value1, value2 string, options DuplicateOptions) DuplicateStatus {
	mu.Lock()
	defer mu.Unlock()

	cValue1 := C.CString(value1)
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := C.CString(value2)
	defer C.free(unsafe.Pointer(cValue2))

	var charPtr *C.char
	ptrSize := unsafe.Sizeof(charPtr)

	cOptions := C.libpostal_get_default_duplicate_options()
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

	return 0
}
