package appindicator

// #include <libappindicator/app-indicator.h>
import "C"

func stringToGCharPointer(s string) *C.gchar {
	return (*C.gchar)(C.CString(s))
}

func uIntToGUInt(i uint) C.guint {
	return (C.guint)(i)
}
