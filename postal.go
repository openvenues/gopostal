package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>

*/
import "C"

import (
	"log"
	"sync"
)

var mu sync.Mutex

func init() {
	if !bool(C.libpostal_setup()) || !bool(C.libpostal_setup_language_classifier()) {
		log.Fatal("Could not load libpostal")
	}
}
