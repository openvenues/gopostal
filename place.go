package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func PlaceLanguages(labels []string, values []string) []string {
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
	cNumLanguages := C.size_t(0)
	cLanguages := C.libpostal_place_languages(cNumComponents, &cLabels[0], &cValues[0], &cNumLanguages)
	cLanguagesPtr := (*[1 << 30](*C.char))(unsafe.Pointer(cLanguages))

	var languages []string
	var i uint64
	for i = 0; i < uint64(cNumLanguages); i++ {
		languages = append(languages, C.GoString(cLanguagesPtr[i]))
	}

	return languages
}
